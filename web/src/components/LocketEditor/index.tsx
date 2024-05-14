import { Button, Divider, IconButton, Option, Select } from "@mui/joy";
import React, { useEffect, useMemo, useRef, useState } from "react";
import { toast } from "react-hot-toast";
import { useTranslation } from "react-i18next";
import useLocalStorage from "react-use/lib/useLocalStorage";
import { locketServiceClient } from "@/grpcweb";
import { TAB_SPACE_WIDTH } from "@/helpers/consts";
import { isValidUrl } from "@/helpers/utils";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useGlobalStore, useResourceStore, useTagStore } from "@/store/module";
import { useLocketStore, useUserStore } from "@/store/v1";
import { LocketRelation, LocketRelation_Type } from "@/types/proto/api/v2/locket_relation_service";
import { Locket, Visibility } from "@/types/proto/api/v2/locket_service";
import { Resource } from "@/types/proto/api/v2/resource_service";
import { UserSetting } from "@/types/proto/api/v2/user_service";
import { useTranslate } from "@/utils/i18n";
import { convertVisibilityFromString, convertVisibilityToString } from "@/utils/locket";
import { extractTagsFromContent } from "@/utils/tag";
import showCreateResourceDialog from "../CreateResourceDialog";
import Icon from "../Icon";
import VisibilityIcon from "../VisibilityIcon";
import AddLocketRelationButton from "./ActionButton/AddLocketRelationButton";
import MarkdownMenu from "./ActionButton/MarkdownMenu";
import TagSelector from "./ActionButton/TagSelector";
import Editor, { EditorRefActions } from "./Editor";
import RelationListView from "./RelationListView";
import ResourceListView from "./ResourceListView";
import { handleEditorKeydownWithMarkdownShortcuts, hyperlinkHighlightedText } from "./handlers";
import { LocketEditorContext } from "./types";

interface Props {
  className?: string;
  cacheKey?: string;
  placeholder?: string;
  locketName?: string;
  parentLocketName?: string;
  relationList?: LocketRelation[];
  autoFocus?: boolean;
  onConfirm?: (locketName: string) => void;
  onEditPrevious?: () => void;
}

interface State {
  locketVisibility: Visibility;
  resourceList: Resource[];
  relationList: LocketRelation[];
  isUploadingResource: boolean;
  isRequesting: boolean;
  isComposing: boolean;
}

const LocketEditor = (props: Props) => {
  const { className, cacheKey, locketName, parentLocketName, autoFocus, onConfirm } = props;
  const { i18n } = useTranslation();
  const t = useTranslate();
  const {
    state: { systemStatus },
  } = useGlobalStore();
  const userStore = useUserStore();
  const locketStore = useLocketStore();
  const resourceStore = useResourceStore();
  const tagStore = useTagStore();
  const currentUser = useCurrentUser();
  const [state, setState] = useState<State>({
    locketVisibility: Visibility.PRIVATE,
    resourceList: [],
    relationList: props.relationList ?? [],
    isUploadingResource: false,
    isRequesting: false,
    isComposing: false,
  });
  const [hasContent, setHasContent] = useState<boolean>(false);
  const editorRef = useRef<EditorRefActions>(null);
  const userSetting = userStore.userSetting as UserSetting;
  const contentCacheKey = `${currentUser.name}-${cacheKey || ""}`;
  const [contentCache, setContentCache] = useLocalStorage<string>(contentCacheKey, "");
  const referenceRelations = locketName
    ? state.relationList.filter(
        (relation) =>
          relation.locket === locketName && relation.relatedLocket !== locketName && relation.type === LocketRelation_Type.REFERENCE,
      )
    : state.relationList.filter((relation) => relation.type === LocketRelation_Type.REFERENCE);

  useEffect(() => {
    editorRef.current?.setContent(contentCache || "");
  }, []);

  useEffect(() => {
    if (autoFocus) {
      handleEditorFocus();
    }
  }, [autoFocus]);

  useEffect(() => {
    let visibility = userSetting.locketVisibility;
    if (systemStatus.disablePublicLockets && visibility === "PUBLIC") {
      visibility = "PRIVATE";
    }
    setState((prevState) => ({
      ...prevState,
      locketVisibility: convertVisibilityFromString(visibility),
    }));
  }, [userSetting.locketVisibility, systemStatus.disablePublicLockets]);

  useEffect(() => {
    if (locketName) {
      locketStore.getOrFetchLocketByName(locketName).then((locket) => {
        if (locket) {
          handleEditorFocus();
          setState((prevState) => ({
            ...prevState,
            locketVisibility: locket.visibility,
            resourceList: locket.resources,
            relationList: locket.relations,
          }));
          if (!contentCache) {
            editorRef.current?.setContent(locket.content ?? "");
          }
        }
      });
    }
  }, [locketName]);

  const handleCompositionStart = () => {
    setState((prevState) => ({
      ...prevState,
      isComposing: true,
    }));
  };

  const handleCompositionEnd = () => {
    setState((prevState) => ({
      ...prevState,
      isComposing: false,
    }));
  };

  const handleKeyDown = (event: React.KeyboardEvent) => {
    if (!editorRef.current) {
      return;
    }

    const isMetaKey = event.ctrlKey || event.metaKey;
    if (isMetaKey) {
      if (event.key === "Enter") {
        void handleSaveBtnClick();
        return;
      }

      handleEditorKeydownWithMarkdownShortcuts(event, editorRef.current);
    }
    if (event.key === "Tab" && !state.isComposing) {
      event.preventDefault();
      const tabSpace = " ".repeat(TAB_SPACE_WIDTH);
      const cursorPosition = editorRef.current.getCursorPosition();
      const selectedContent = editorRef.current.getSelectedContent();
      editorRef.current.insertText(tabSpace);
      if (selectedContent) {
        editorRef.current.setCursorPosition(cursorPosition + TAB_SPACE_WIDTH);
      }
      return;
    }

    if (!!props.onEditPrevious && event.key === "ArrowDown" && !state.isComposing && editorRef.current.getContent() === "") {
      event.preventDefault();
      props.onEditPrevious();
      return;
    }
  };

  const handleLocketVisibilityChange = (visibility: Visibility) => {
    setState((prevState) => ({
      ...prevState,
      locketVisibility: visibility,
    }));
  };

  const handleUploadFileBtnClick = () => {
    showCreateResourceDialog({
      onConfirm: (resourceList) => {
        setState((prevState) => ({
          ...prevState,
          resourceList: [...prevState.resourceList, ...resourceList],
        }));
      },
    });
  };

  const handleSetResourceList = (resourceList: Resource[]) => {
    setState((prevState) => ({
      ...prevState,
      resourceList,
    }));
  };

  const handleSetRelationList = (relationList: LocketRelation[]) => {
    setState((prevState) => ({
      ...prevState,
      relationList,
    }));
  };

  const handleUploadResource = async (file: File) => {
    setState((state) => {
      return {
        ...state,
        isUploadingResource: true,
      };
    });

    let resource = undefined;
    try {
      resource = await resourceStore.createResourceWithBlob(file);
    } catch (error: any) {
      console.error(error);
      toast.error(typeof error === "string" ? error : error.response.data.message);
    }

    setState((state) => {
      return {
        ...state,
        isUploadingResource: false,
      };
    });
    return resource;
  };

  const uploadMultiFiles = async (files: FileList) => {
    const uploadedResourceList: Resource[] = [];
    for (const file of files) {
      const resource = await handleUploadResource(file);
      if (resource) {
        uploadedResourceList.push(resource);
        if (locketName) {
          await resourceStore.updateResource({
            resource: Resource.fromPartial({
              name: resource.name,
              locket: locketName,
            }),
            updateMask: ["locket"],
          });
        }
      }
    }
    if (uploadedResourceList.length > 0) {
      setState((prevState) => ({
        ...prevState,
        resourceList: [...prevState.resourceList, ...uploadedResourceList],
      }));
    }
  };

  const handleDropEvent = async (event: React.DragEvent) => {
    if (event.dataTransfer && event.dataTransfer.files.length > 0) {
      event.preventDefault();
      await uploadMultiFiles(event.dataTransfer.files);
    }
  };

  const handlePasteEvent = async (event: React.ClipboardEvent) => {
    if (event.clipboardData && event.clipboardData.files.length > 0) {
      event.preventDefault();
      await uploadMultiFiles(event.clipboardData.files);
    } else if (
      editorRef.current != null &&
      editorRef.current.getSelectedContent().length != 0 &&
      isValidUrl(event.clipboardData.getData("Text"))
    ) {
      event.preventDefault();
      hyperlinkHighlightedText(editorRef.current, event.clipboardData.getData("Text"));
    }
  };

  const handleContentChange = (content: string) => {
    setHasContent(content !== "");
    if (content !== "") {
      setContentCache(content);
    } else {
      localStorage.removeItem(contentCacheKey);
    }
  };

  const handleSaveBtnClick = async () => {
    if (state.isRequesting) {
      return;
    }

    setState((state) => {
      return {
        ...state,
        isRequesting: true,
      };
    });
    const content = editorRef.current?.getContent() ?? "";
    try {
      // Update locket.
      if (locketName) {
        const prevLocket = await locketStore.getOrFetchLocketByName(locketName);
        if (prevLocket) {
          const locket = await locketStore.updateLocket(
            {
              name: prevLocket.name,
              content,
              visibility: state.locketVisibility,
            },
            ["content", "visibility"],
          );
          await locketServiceClient.setLocketResources({
            name: locket.name,
            resources: state.resourceList,
          });
          await locketServiceClient.setLocketRelations({
            name: locket.name,
            relations: state.relationList,
          });
          await locketStore.getOrFetchLocketByName(locket.name, { skipCache: true });
          if (onConfirm) {
            onConfirm(locket.name);
          }
        }
      } else {
        // Create locket or locket comment.
        const request = !parentLocketName
          ? locketStore.createLocket({
              content,
              visibility: state.locketVisibility,
            })
          : locketServiceClient
              .createLocketComment({
                name: parentLocketName,
                comment: {
                  content,
                  visibility: state.locketVisibility,
                },
              })
              .then(({ locket }) => locket as Locket);
        const locket = await request;
        await locketServiceClient.setLocketResources({
          name: locket.name,
          resources: state.resourceList,
        });
        await locketServiceClient.setLocketRelations({
          name: locket.name,
          relations: state.relationList,
        });
        await locketStore.getOrFetchLocketByName(locket.name, { skipCache: true });
        if (onConfirm) {
          onConfirm(locket.name);
        }
      }
      editorRef.current?.setContent("");
    } catch (error: any) {
      console.error(error);
      toast.error(error.details);
    }

    // Batch upsert tags.
    const tags = extractTagsFromContent(content);
    await tagStore.batchUpsertTag(tags);

    setState((state) => {
      return {
        ...state,
        isRequesting: false,
        resourceList: [],
        relationList: [],
      };
    });
  };

  const handleEditorFocus = () => {
    editorRef.current?.focus();
  };

  const editorConfig = useMemo(
    () => ({
      className: "",
      initialContent: "",
      placeholder: props.placeholder ?? t("editor.any-thoughts"),
      onContentChange: handleContentChange,
      onPaste: handlePasteEvent,
    }),
    [i18n.language],
  );

  const allowSave = (hasContent || state.resourceList.length > 0) && !state.isUploadingResource && !state.isRequesting;

  return (
    <LocketEditorContext.Provider
      value={{
        relationList: state.relationList,
        setRelationList: (relationList: LocketRelation[]) => {
          setState((prevState) => ({
            ...prevState,
            relationList,
          }));
        },
        locketName,
      }}
    >
      <div
        className={`${
          className ?? ""
        } relative w-full flex flex-col justify-start items-start bg-white dark:bg-zinc-800 px-4 pt-4 rounded-lg border border-gray-200 dark:border-zinc-700`}
        tabIndex={0}
        onKeyDown={handleKeyDown}
        onDrop={handleDropEvent}
        onFocus={handleEditorFocus}
        onCompositionStart={handleCompositionStart}
        onCompositionEnd={handleCompositionEnd}
      >
        <Editor ref={editorRef} {...editorConfig} />
        <ResourceListView resourceList={state.resourceList} setResourceList={handleSetResourceList} />
        <RelationListView relationList={referenceRelations} setRelationList={handleSetRelationList} />
        <div className="relative w-full flex flex-row justify-between items-center pt-2" onFocus={(e) => e.stopPropagation()}>
          <div className="flex flex-row justify-start items-center opacity-80">
            <TagSelector editorRef={editorRef} />
            <MarkdownMenu editorRef={editorRef} />
            <IconButton size="sm" onClick={handleUploadFileBtnClick}>
              <Icon.Image className="w-5 h-5 mx-auto" />
            </IconButton>
            <AddLocketRelationButton editorRef={editorRef} />
          </div>
        </div>
        <Divider className="!mt-2" />
        <div className="w-full flex flex-row justify-between items-center py-3 dark:border-t-zinc-500">
          <div className="relative flex flex-row justify-start items-center" onFocus={(e) => e.stopPropagation()}>
            <Select
              variant="plain"
              value={state.locketVisibility}
              startDecorator={<VisibilityIcon visibility={state.locketVisibility} />}
              onChange={(_, visibility) => {
                if (visibility) {
                  handleLocketVisibilityChange(visibility);
                }
              }}
            >
              {[Visibility.PRIVATE, Visibility.PROTECTED, Visibility.PUBLIC].map((item) => (
                <Option key={item} value={item} className="whitespace-nowrap">
                  {t(`locket.visibility.${convertVisibilityToString(item).toLowerCase()}` as any)}
                </Option>
              ))}
            </Select>
          </div>
          <div className="shrink-0 flex flex-row justify-end items-center">
            <Button
              disabled={!allowSave}
              loading={state.isRequesting}
              endDecorator={<Icon.Send className="w-4 h-auto" />}
              onClick={handleSaveBtnClick}
            >
              {t("editor.save")}
            </Button>
          </div>
        </div>
      </div>
    </LocketEditorContext.Provider>
  );
};

export default LocketEditor;
