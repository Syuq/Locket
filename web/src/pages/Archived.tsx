import { Button, Tooltip } from "@mui/joy";
import { ClientError } from "nice-grpc-web";
import { useEffect, useRef, useState } from "react";
import toast from "react-hot-toast";
import { showCommonDialog } from "@/components/Dialog/CommonDialog";
import Empty from "@/components/Empty";
import Icon from "@/components/Icon";
import LocketContent from "@/components/LocketContent";
import LocketFilter from "@/components/LocketFilter";
import MobileHeader from "@/components/MobileHeader";
import SearchBar from "@/components/SearchBar";
import { DEFAULT_LIST_LOCKETS_PAGE_SIZE } from "@/helpers/consts";
import { getTimeStampByDate } from "@/helpers/datetime";
import useCurrentUser from "@/hooks/useCurrentUser";
import useFilterWithUrlParams from "@/hooks/useFilterWithUrlParams";
import { useLocketList, useLocketStore } from "@/store/v1";
import { RowStatus } from "@/types/proto/api/v2/common";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { useTranslate } from "@/utils/i18n";

const Archived = () => {
  const t = useTranslate();
  const user = useCurrentUser();
  const locketStore = useLocketStore();
  const locketList = useLocketList();
  const [isRequesting, setIsRequesting] = useState(true);
  const nextPageTokenRef = useRef<string | undefined>(undefined);
  const { tag: tagQuery, text: textQuery } = useFilterWithUrlParams();
  const sortedLockets = locketList.value
    .filter((locket) => locket.rowStatus === RowStatus.ARCHIVED)
    .sort((a, b) => getTimeStampByDate(b.displayTime) - getTimeStampByDate(a.displayTime));

  useEffect(() => {
    nextPageTokenRef.current = undefined;
    locketList.reset();
    fetchLockets();
  }, [tagQuery, textQuery]);

  const fetchLockets = async () => {
    const filters = [`creator == "${user.name}"`, `row_status == "ARCHIVED"`];
    const contentSearch: string[] = [];
    if (tagQuery) {
      contentSearch.push(JSON.stringify(`#${tagQuery}`));
    }
    if (textQuery) {
      contentSearch.push(JSON.stringify(textQuery));
    }
    if (contentSearch.length > 0) {
      filters.push(`content_search == [${contentSearch.join(", ")}]`);
    }
    setIsRequesting(true);
    const data = await locketStore.fetchLockets({
      pageSize: DEFAULT_LIST_LOCKETS_PAGE_SIZE,
      filter: filters.join(" && "),
      pageToken: nextPageTokenRef.current,
    });
    setIsRequesting(false);
    nextPageTokenRef.current = data.nextPageToken;
  };

  const handleDeleteLocketClick = async (locket: Locket) => {
    showCommonDialog({
      title: t("locket.delete-locket"),
      content: t("locket.delete-confirm"),
      style: "danger",
      dialogName: "delete-locket-dialog",
      onConfirm: async () => {
        await locketStore.deleteLocket(locket.name);
      },
    });
  };

  const handleRestoreLocketClick = async (locket: Locket) => {
    try {
      await locketStore.updateLocket(
        {
          name: locket.name,
          rowStatus: RowStatus.ACTIVE,
        },
        ["row_status"],
      );
      toast(t("message.restored-successfully"));
    } catch (error: unknown) {
      console.error(error);
      toast.error((error as ClientError).details);
    }
  };

  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      <MobileHeader />
      <div className="w-full px-4 sm:px-6">
        <div className="w-full flex flex-col justify-start items-start">
          <div className="w-full flex flex-row justify-end items-center mb-2">
            <div className="w-40">
              <SearchBar />
            </div>
          </div>
          <LocketFilter className="px-2 pb-2" />
          {sortedLockets.map((locket) => (
            <div
              key={locket.name}
              className="relative flex flex-col justify-start items-start w-full p-4 pt-3 mb-2 bg-white dark:bg-zinc-800 rounded-lg"
            >
              <div className="w-full mb-1 flex flex-row justify-between items-center">
                <div className="w-full max-w-[calc(100%-20px)] flex flex-row justify-start items-center mr-1">
                  <div className="text-sm leading-6 text-gray-400 select-none">
                    <relative-time datetime={locket.displayTime?.toISOString()} tense="past"></relative-time>
                  </div>
                </div>
                <div className="flex flex-row justify-end items-center gap-x-2">
                  <Tooltip title={t("common.restore")} placement="top">
                    <button onClick={() => handleRestoreLocketClick(locket)}>
                      <Icon.ArchiveRestore className="w-4 h-auto cursor-pointer text-gray-500 dark:text-gray-400" />
                    </button>
                  </Tooltip>
                  <Tooltip title={t("common.delete")} placement="top">
                    <button onClick={() => handleDeleteLocketClick(locket)} className="text-gray-500 dark:text-gray-400">
                      <Icon.Trash className="w-4 h-auto cursor-pointer" />
                    </button>
                  </Tooltip>
                </div>
              </div>
              <LocketContent
                key={`${locket.name}-${locket.displayTime}`}
                locketName={locket.name}
                content={locket.content}
                readonly={true}
              />
            </div>
          ))}
          {isRequesting ? (
            <div className="flex flex-row justify-center items-center w-full my-4 text-gray-400">
              <Icon.Loader className="w-4 h-auto animate-spin mr-1" />
              <p className="text-sm italic">{t("locket.fetching-data")}</p>
            </div>
          ) : !nextPageTokenRef.current ? (
            sortedLockets.length === 0 && (
              <div className="w-full mt-16 mb-8 flex flex-col justify-center items-center italic">
                <Empty />
                <p className="mt-4 text-gray-600 dark:text-gray-400">{t("message.no-data")}</p>
              </div>
            )
          ) : (
            <div className="w-full flex flex-row justify-center items-center my-4">
              <Button variant="plain" endDecorator={<Icon.ArrowDown className="w-5 h-auto" />} onClick={fetchLockets}>
                {t("locket.fetch-more")}
              </Button>
            </div>
          )}
        </div>
      </div>
    </section>
  );
};

export default Archived;
