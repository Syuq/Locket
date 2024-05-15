import { Tooltip } from "@mui/joy";
import classNames from "classnames";
import { memo, useCallback, useEffect, useRef, useState } from "react";
import { Link, useLocation } from "react-router-dom";
import useCurrentUser from "@/hooks/useCurrentUser";
import useNavigateTo from "@/hooks/useNavigateTo";
import { extractLocketIdFromName, useUserStore } from "@/store/v1";
import { LocketRelation_Type } from "@/types/proto/api/v2/locket_relation_service";
import { Locket, Visibility } from "@/types/proto/api/v2/locket_service";
import { useTranslate } from "@/utils/i18n";
import { convertVisibilityToString } from "@/utils/locket";
import showChangeLocketCreatedTsDialog from "./ChangeLocketCreatedTsDialog";
import Icon from "./Icon";
import LocketActionMenu from "./LocketActionMenu";
import LocketContent from "./LocketContent";
import LocketReactionistView from "./LocketReactionListView";
import LocketRelationListView from "./LocketRelationListView";
import LocketResourceListView from "./LocketResourceListView";
import showPreviewImageDialog from "./PreviewImageDialog";
import ReactionSelector from "./ReactionSelector";
import UserAvatar from "./UserAvatar";
import VisibilityIcon from "./VisibilityIcon";

interface Props {
  locket: Locket;
  compact?: boolean;
  showCreator?: boolean;
  showVisibility?: boolean;
  showPinned?: boolean;
  className?: string;
}

const LocketView: React.FC<Props> = (props: Props) => {
  const { locket, className } = props;
  const t = useTranslate();
  const location = useLocation();
  const navigateTo = useNavigateTo();
  const currentUser = useCurrentUser();
  const userStore = useUserStore();
  const user = useCurrentUser();
  const [creator, setCreator] = useState(userStore.getUserByName(locket.creator));
  const locketContainerRef = useRef<HTMLDivElement>(null);
  const referencedLockets = locket.relations.filter((relation) => relation.type === LocketRelation_Type.REFERENCE);
  const commentAmount = locket.relations.filter(
    (relation) => relation.type === LocketRelation_Type.COMMENT && relation.relatedLocket === locket.name,
  ).length;
  const relativeTimeFormat = Date.now() - locket.displayTime!.getTime() > 1000 * 60 * 60 * 24 ? "datetime" : "auto";
  const readonly = locket.creator !== user?.name;
  const isInLocketDetailPage = location.pathname.startsWith(`/m/${locket.uid}`);

  // Initial related data: creator.
  useEffect(() => {
    (async () => {
      const user = await userStore.getOrFetchUserByName(locket.creator);
      setCreator(user);
    })();
  }, []);

  const handleGotoLocketDetailPage = (event: React.MouseEvent<HTMLDivElement>) => {
    if (event.altKey) {
      showChangeLocketCreatedTsDialog(extractLocketIdFromName(locket.name));
    } else {
      navigateTo(`/m/${locket.uid}`);
    }
  };

  const handleLocketContentClick = useCallback(async (e: React.MouseEvent) => {
    const targetEl = e.target as HTMLElement;

    if (targetEl.tagName === "IMG") {
      const imgUrl = targetEl.getAttribute("src");
      if (imgUrl) {
        showPreviewImageDialog([imgUrl], 0);
      }
    }
  }, []);

  return (
    <div
      className={classNames(
        "group relative flex flex-col justify-start items-start w-full px-4 py-3 mb-2 gap-2 bg-white dark:bg-zinc-800 rounded-lg border border-white dark:border-zinc-800 hover:border-gray-200 dark:hover:border-zinc-700",
        props.showPinned && locket.pinned && "border-gray-200 border dark:border-zinc-700",
        className,
      )}
      ref={locketContainerRef}
    >
      <div className="w-full flex flex-row justify-between items-center gap-2">
        <div className="w-auto max-w-[calc(100%-8rem)] grow flex flex-row justify-start items-center">
          {props.showCreator && creator ? (
            <div className="w-full flex flex-row justify-start items-center">
              <Link className="w-auto hover:opacity-80" to={`/u/${encodeURIComponent(creator.username)}`} unstable_viewTransition>
                <UserAvatar className="mr-2 shrink-0" avatarUrl={creator.avatarUrl} />
              </Link>
              <div className="w-full flex flex-col justify-center items-start">
                <Link
                  className="w-full block leading-tight hover:opacity-80 truncate text-gray-600 dark:text-gray-400"
                  to={`/u/${encodeURIComponent(creator.username)}`}
                  unstable_viewTransition
                >
                  {creator.nickname || creator.username}
                </Link>
                <div
                  className="w-auto -mt-0.5 text-xs leading-tight text-gray-400 dark:text-gray-500 select-none"
                  onClick={handleGotoLocketDetailPage}
                >
                  <relative-time datetime={locket.displayTime?.toISOString()} format={relativeTimeFormat} tense="past"></relative-time>
                </div>
              </div>
            </div>
          ) : (
            <div className="w-full text-sm leading-tight text-gray-400 dark:text-gray-500 select-none" onClick={handleGotoLocketDetailPage}>
              <relative-time datetime={locket.displayTime?.toISOString()} format={relativeTimeFormat} tense="past"></relative-time>
            </div>
          )}
        </div>
        <div className="flex flex-row justify-end items-center select-none shrink-0 gap-2">
          <div className="w-auto invisible group-hover:visible flex flex-row justify-between items-center gap-2">
            {props.showVisibility && locket.visibility !== Visibility.PRIVATE && (
              <Tooltip title={t(`locket.visibility.${convertVisibilityToString(locket.visibility).toLowerCase()}` as any)} placement="top">
                <span className="flex justify-center items-center hover:opacity-70">
                  <VisibilityIcon visibility={locket.visibility} />
                </span>
              </Tooltip>
            )}
            {currentUser && <ReactionSelector className="border-none w-auto h-auto" locket={locket} />}
          </div>
          {!isInLocketDetailPage && (
            <Link
              className={classNames(
                "flex flex-row justify-start items-center hover:opacity-70",
                commentAmount === 0 && "invisible group-hover:visible",
              )}
              to={`/m/${locket.uid}#comments`}
              unstable_viewTransition
            >
              <Icon.MessageCircleMore className="w-4 h-4 mx-auto text-gray-500 dark:text-gray-400" />
              {commentAmount > 0 && <span className="text-xs text-gray-500 dark:text-gray-400">{commentAmount}</span>}
            </Link>
          )}
          {props.showPinned && locket.pinned && (
            <Tooltip title={"Pinned"} placement="top">
              <Icon.Bookmark className="w-4 h-auto text-amber-500" />
            </Tooltip>
          )}
          {!readonly && <LocketActionMenu className="-ml-1" locket={locket} hiddenActions={props.showPinned ? [] : ["pin"]} />}
        </div>
      </div>
      <LocketContent
        key={`${locket.name}-${locket.updateTime}`}
        locketName={locket.name}
        content={locket.content}
        readonly={readonly}
        onClick={handleLocketContentClick}
        compact={props.compact ?? true}
      />
      <LocketResourceListView resources={locket.resources} />
      <LocketRelationListView locket={locket} relations={referencedLockets} />
      <LocketReactionistView locket={locket} reactions={locket.reactions} />
    </div>
  );
};

export default memo(LocketView);
