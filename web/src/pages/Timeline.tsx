import { Button, Divider, IconButton } from "@mui/joy";
import classNames from "classnames";
import { Fragment, useEffect, useRef, useState } from "react";
import ActivityCalendar from "@/components/ActivityCalendar";
import Empty from "@/components/Empty";
import Icon from "@/components/Icon";
import showLocketEditorDialog from "@/components/LocketEditor/LocketEditorDialog";
import LocketFilter from "@/components/LocketFilter";
import LocketView from "@/components/LocketView";
import MobileHeader from "@/components/MobileHeader";
import { TimelineSidebar, TimelineSidebarDrawer } from "@/components/TimelineSidebar";
import { locketServiceClient } from "@/grpcweb";
import { DAILY_TIMESTAMP, DEFAULT_LIST_LOCKETS_PAGE_SIZE } from "@/helpers/consts";
import { getNormalizedTimeString, getTimeStampByDate } from "@/helpers/datetime";
import useCurrentUser from "@/hooks/useCurrentUser";
import useFilterWithUrlParams from "@/hooks/useFilterWithUrlParams";
import useResponsiveWidth from "@/hooks/useResponsiveWidth";
import i18n from "@/i18n";
import { useLocketList, useLocketStore } from "@/store/v1";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { useTranslate } from "@/utils/i18n";

interface GroupedByMonthItem {
  // Format: 2021-1
  month: string;
  data: Record<string, number>;
  lockets: Locket[];
}

const groupByMonth = (dateCountMap: Record<string, number>, lockets: Locket[]): GroupedByMonthItem[] => {
  const groupedByMonth: GroupedByMonthItem[] = [];

  Object.entries(dateCountMap).forEach(([date, count]) => {
    const month = date.split("-").slice(0, 2).join("-");
    const existingMonth = groupedByMonth.find((group) => group.month === month);
    if (existingMonth) {
      existingMonth.data[date] = count;
    } else {
      const monthLockets = lockets.filter((locket) => getNormalizedTimeString(locket.displayTime).startsWith(month));
      groupedByMonth.push({ month, data: { [date]: count }, lockets: monthLockets });
    }
  });

  return groupedByMonth
    .filter((group) => group.lockets.length > 0)
    .sort((a, b) => getTimeStampByDate(b.month) - getTimeStampByDate(a.month));
};

const Timeline = () => {
  const t = useTranslate();
  const { md } = useResponsiveWidth();
  const user = useCurrentUser();
  const locketStore = useLocketStore();
  const locketList = useLocketList();
  const [activityStats, setActivityStats] = useState<Record<string, number>>({});
  const [selectedDay, setSelectedDay] = useState<string | undefined>();
  const [isRequesting, setIsRequesting] = useState(true);
  const nextPageTokenRef = useRef<string | undefined>(undefined);
  const { tag: tagQuery, text: textQuery } = useFilterWithUrlParams();
  const sortedLockets = locketList.value.sort((a, b) => getTimeStampByDate(b.displayTime) - getTimeStampByDate(a.displayTime));
  const groupedByMonth = groupByMonth(activityStats, sortedLockets);

  useEffect(() => {
    nextPageTokenRef.current = undefined;
    locketList.reset();
    fetchLockets();
  }, [selectedDay, tagQuery, textQuery]);

  useEffect(() => {
    (async () => {
      const filters = [`row_status == "NORMAL"`];
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
      const { stats } = await locketServiceClient.getUserLocketsStats({
        name: user.name,
        timezone: Intl.DateTimeFormat().resolvedOptions().timeZone,
        filter: filters.join(" && "),
      });
      setActivityStats(stats);
    })();
  }, [sortedLockets.length]);

  const fetchLockets = async () => {
    const filters = [`creator == "${user.name}"`, `row_status == "NORMAL"`];
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
    if (selectedDay) {
      const selectedDateStamp = getTimeStampByDate(selectedDay) + new Date().getTimezoneOffset() * 60 * 1000;
      filters.push(
        ...[`display_time_after == ${selectedDateStamp / 1000}`, `display_time_before == ${(selectedDateStamp + DAILY_TIMESTAMP) / 1000}`],
      );
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

  const handleNewLocket = () => {
    showLocketEditorDialog({});
  };

  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      {!md && (
        <MobileHeader>
          <TimelineSidebarDrawer />
        </MobileHeader>
      )}
      <div className={classNames("w-full flex flex-row justify-start items-start px-4 sm:px-6 gap-4")}>
        <div className={classNames(md ? "w-[calc(100%-15rem)]" : "w-full")}>
          <div className="w-full shadow flex flex-col justify-start items-start px-4 py-3 rounded-xl bg-white dark:bg-zinc-800 text-black dark:text-gray-300">
            <div className="relative w-full flex flex-row justify-between items-center">
              <div>
                <div
                  className="py-1 flex flex-row justify-start items-center select-none opacity-80"
                  onClick={() => setSelectedDay(undefined)}
                >
                  <Icon.GanttChartSquare className="w-6 h-auto mr-1 opacity-80" />
                  <span className="text-lg">{t("timeline.title")}</span>
                </div>
              </div>
              <div className="flex justify-end items-center gap-2">
                <IconButton variant="outlined" size="sm" onClick={() => handleNewLocket()}>
                  <Icon.Plus className="w-5 h-auto" />
                </IconButton>
              </div>
            </div>
            <div className="w-full h-auto flex flex-col justify-start items-start">
              <LocketFilter className="px-2 my-4" />

              {groupedByMonth.map((group, index) => (
                <Fragment key={group.month}>
                  <div className={classNames("flex flex-col justify-start items-start w-full mt-2 last:mb-4")}>
                    <div className={classNames("flex shrink-0 flex-row w-full pl-1 mt-2 mb-2")}>
                      <div className={classNames("w-full flex flex-col")}>
                        <span className="font-medium text-3xl leading-tight mb-1">
                          {new Date(group.month).toLocaleString(i18n.language, { month: "short", timeZone: "UTC" })}
                        </span>
                        <span className="opacity-60">{new Date(group.month).getUTCFullYear()}</span>
                      </div>
                      <ActivityCalendar month={group.month} data={group.data} onClick={(date) => setSelectedDay(date)} />
                    </div>

                    <div className={classNames("w-full flex flex-col justify-start items-start")}>
                      {group.lockets.map((locket, index) => (
                        <div
                          key={`${locket.name}-${locket.displayTime}`}
                          className={classNames("relative w-full flex flex-col justify-start items-start pl-4 sm:pl-10 pt-0")}
                        >
                          <LocketView className="!border max-w-full !border-gray-100 dark:!border-zinc-700" locket={locket} />
                          <div className="absolute -left-2 sm:left-2 top-4 h-full">
                            {index !== group.lockets.length - 1 && (
                              <div className="absolute top-2 left-[7px] h-full w-0.5 bg-gray-200 dark:bg-gray-700 block"></div>
                            )}
                            <div className="border-4 rounded-full border-white relative dark:border-zinc-800">
                              <Icon.Circle className="w-2 h-auto bg-gray-200 text-gray-200 dark:bg-gray-700 dark:text-gray-700 rounded-full" />
                            </div>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>
                  {index !== groupedByMonth.length - 1 && <Divider className="w-full !my-4 md:!mb-8 !bg-gray-100 dark:!bg-zinc-700" />}
                </Fragment>
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
        </div>
        {md && (
          <div className="sticky top-0 left-0 shrink-0 -mt-6 w-56 h-full">
            <TimelineSidebar className="py-6" />
          </div>
        )}
      </div>
    </section>
  );
};

export default Timeline;
