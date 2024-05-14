import { Button } from "@mui/joy";
import classNames from "classnames";
import { useEffect, useRef, useState } from "react";
import Empty from "@/components/Empty";
import { ExploreSidebar, ExploreSidebarDrawer } from "@/components/ExploreSidebar";
import Icon from "@/components/Icon";
import LocketFilter from "@/components/LocketFilter";
import LocketView from "@/components/LocketView";
import MobileHeader from "@/components/MobileHeader";
import { DEFAULT_LIST_LOCKETS_PAGE_SIZE } from "@/helpers/consts";
import { getTimeStampByDate } from "@/helpers/datetime";
import useCurrentUser from "@/hooks/useCurrentUser";
import useFilterWithUrlParams from "@/hooks/useFilterWithUrlParams";
import useResponsiveWidth from "@/hooks/useResponsiveWidth";
import { useLocketList, useLocketStore } from "@/store/v1";
import { useTranslate } from "@/utils/i18n";

const Explore = () => {
  const t = useTranslate();
  const { md } = useResponsiveWidth();
  const user = useCurrentUser();
  const locketStore = useLocketStore();
  const locketList = useLocketList();
  const [isRequesting, setIsRequesting] = useState(true);
  const nextPageTokenRef = useRef<string | undefined>(undefined);
  const { tag: tagQuery, text: textQuery } = useFilterWithUrlParams();
  const sortedLockets = locketList.value.sort((a, b) => getTimeStampByDate(b.displayTime) - getTimeStampByDate(a.displayTime));

  useEffect(() => {
    nextPageTokenRef.current = undefined;
    locketList.reset();
    fetchLockets();
  }, [tagQuery, textQuery]);

  const fetchLockets = async () => {
    const filters = [`row_status == "NORMAL"`, `visibilities == [${user ? "'PUBLIC', 'PROTECTED'" : "'PUBLIC'"}]`];
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

  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      {!md && (
        <MobileHeader>
          <ExploreSidebarDrawer />
        </MobileHeader>
      )}
      <div className={classNames("w-full flex flex-row justify-start items-start px-4 sm:px-6 gap-4")}>
        <div className={classNames(md ? "w-[calc(100%-15rem)]" : "w-full")}>
          <div className="flex flex-col justify-start items-start w-full max-w-full">
            <LocketFilter className="px-2 pb-2" />
            {sortedLockets.map((locket) => (
              <LocketView key={`${locket.name}-${locket.updateTime}`} locket={locket} showCreator showVisibility showPinned />
            ))}
            {isRequesting ? (
              <div className="flex flex-row justify-center items-center w-full my-4 text-gray-400">
                <Icon.Loader className="w-4 h-auto animate-spin mr-1" />
                <p className="text-sm italic">{t("locket.fetching-data")}</p>
              </div>
            ) : !nextPageTokenRef.current ? (
              sortedLockets.length === 0 && (
                <div className="w-full mt-12 mb-8 flex flex-col justify-center items-center italic">
                  <Empty />
                  <p className="mt-2 text-gray-600 dark:text-gray-400">{t("message.no-data")}</p>
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
        {md && (
          <div className="sticky top-0 left-0 shrink-0 -mt-6 w-56 h-full">
            <ExploreSidebar className="py-6" />
          </div>
        )}
      </div>
    </section>
  );
};

export default Explore;
