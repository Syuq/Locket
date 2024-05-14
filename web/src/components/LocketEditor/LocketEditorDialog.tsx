import { IconButton } from "@mui/joy";
import { useEffect } from "react";
import { useGlobalStore, useTagStore } from "@/store/module";
import { LocketRelation } from "@/types/proto/api/v2/locket_relation_service";
import LocketEditor from ".";
import { generateDialog } from "../Dialog";
import Icon from "../Icon";

interface Props extends DialogProps {
  locketName?: string;
  cacheKey?: string;
  relationList?: LocketRelation[];
}

const LocketEditorDialog: React.FC<Props> = ({ locketName: locket, cacheKey, relationList, destroy }: Props) => {
  const globalStore = useGlobalStore();
  const tagStore = useTagStore();
  const { systemStatus } = globalStore.state;

  useEffect(() => {
    tagStore.fetchTags();
  }, []);

  const handleCloseBtnClick = () => {
    destroy();
  };

  return (
    <>
      <div className="w-full flex flex-row justify-between items-center mb-2">
        <div className="flex flex-row justify-start items-center">
          <img className="w-6 h-auto rounded-full shadow" src={systemStatus.customizedProfile.logoUrl || "/full-logo.webp"} alt="" />
          <p className="ml-1 text-lg opacity-80 dark:text-gray-300">{systemStatus.customizedProfile.name}</p>
        </div>
        <IconButton size="sm" onClick={handleCloseBtnClick}>
          <Icon.X className="w-5 h-auto" />
        </IconButton>
      </div>
      <div className="flex flex-col justify-start items-start max-w-full w-[36rem]">
        <LocketEditor
          className="border-none !p-0 -mb-2"
          cacheKey={`locket-editor-${cacheKey || locket}`}
          locketName={locket}
          relationList={relationList}
          onConfirm={handleCloseBtnClick}
          autoFocus
        />
      </div>
    </>
  );
};

export default function showLocketEditorDialog(props: Pick<Props, "locketName" | "cacheKey" | "relationList"> = {}): void {
  generateDialog(
    {
      className: "locket-editor-dialog",
      dialogName: "locket-editor-dialog",
    },
    LocketEditorDialog,
    props,
  );
}
