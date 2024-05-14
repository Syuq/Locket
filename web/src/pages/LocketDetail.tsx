import { ClientError } from "nice-grpc-web";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { Link, useParams } from "react-router-dom";
import Icon from "@/components/Icon";
import LocketEditor from "@/components/LocketEditor";
import LocketView from "@/components/LocketView";
import MobileHeader from "@/components/MobileHeader";
import useCurrentUser from "@/hooks/useCurrentUser";
import useNavigateTo from "@/hooks/useNavigateTo";
import { LocketNamePrefix, useLocketStore } from "@/store/v1";
import { LocketRelation_Type } from "@/types/proto/api/v2/locket_relation_service";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { useTranslate } from "@/utils/i18n";

const LocketDetail = () => {
  const t = useTranslate();
  const params = useParams();
  const navigateTo = useNavigateTo();
  const currentUser = useCurrentUser();
  const locketStore = useLocketStore();
  const uid = params.uid;
  const locket = locketStore.getLocketByUid(uid || "");
  const [parentLocket, setParentLocket] = useState<Locket | undefined>(undefined);
  const commentRelations =
    locket?.relations.filter((relation) => relation.relatedLocket === locket.name && relation.type === LocketRelation_Type.COMMENT) || [];
  const comments = commentRelations
    .map((relation) => locketStore.getLocketByName(relation.locket))
    .filter((locket) => locket) as any as Locket[];

  // Prepare locket.
  useEffect(() => {
    if (uid) {
      locketStore.searchLockets(`uid == "${uid}"`).catch((error: ClientError) => {
        toast.error(error.details);
        navigateTo("/403");
      });
    } else {
      navigateTo("/404");
    }
  }, [uid]);

  // Prepare locket comments.
  useEffect(() => {
    if (!locket) {
      return;
    }

    (async () => {
      if (locket.parentId) {
        locketStore.getOrFetchLocketByName(`${LocketNamePrefix}${locket.parentId}`).then((locket: Locket) => {
          setParentLocket(locket);
        });
      } else {
        setParentLocket(undefined);
      }
      await Promise.all(commentRelations.map((relation) => locketStore.getOrFetchLocketByName(relation.locket)));
    })();
  }, [locket]);

  if (!locket) {
    return null;
  }

  const handleCommentCreated = async (locketCommentName: string) => {
    await locketStore.getOrFetchLocketByName(locketCommentName);
    await locketStore.getOrFetchLocketByName(locket.name, { skipCache: true });
  };

  return (
    <section className="@container w-full max-w-5xl min-h-full flex flex-col justify-start items-center sm:pt-3 md:pt-6 pb-8">
      <MobileHeader />
      <div className="w-full px-4 sm:px-6">
        {parentLocket && (
          <div className="w-auto inline-block mb-2">
            <Link
              className="px-3 py-1 border rounded-lg max-w-xs w-auto text-sm flex flex-row justify-start items-center flex-nowrap text-gray-600 dark:text-gray-400 dark:border-gray-500 hover:shadow hover:opacity-80"
              to={`/m/${parentLocket.uid}`}
              unstable_viewTransition
            >
              <Icon.ArrowUpLeftFromCircle className="w-4 h-auto shrink-0 opacity-60 mr-2" />
              <span className="truncate">{parentLocket.content}</span>
            </Link>
          </div>
        )}
        <LocketView
          key={`${locket.name}-${locket.displayTime}`}
          className="shadow hover:shadow-xl transition-all"
          locket={locket}
          compact={false}
          showCreator
          showVisibility
          showPinned
        />
        <div className="pt-8 pb-16 w-full">
          <h2 id="comments" className="sr-only">
            Comments
          </h2>
          <div className="relative mx-auto flex-grow w-full min-h-full flex flex-col justify-start items-start gap-y-1">
            {comments.length === 0 ? (
              <div className="w-full flex flex-col justify-center items-center py-6 mb-2">
                <Icon.MessageCircle strokeWidth={1} className="w-8 h-auto text-gray-400" />
                <p className="text-gray-400 italic text-sm">{t("locket.comment.no-comment")}</p>
              </div>
            ) : (
              <>
                <div className="w-full flex flex-row justify-start items-center pl-3 mb-3">
                  <Icon.MessageCircle className="w-5 h-auto text-gray-400 mr-1" />
                  <span className="text-gray-400 text-sm">{t("locket.comment.self")}</span>
                  <span className="text-gray-400 text-sm ml-0.5">({comments.length})</span>
                </div>
                {comments.map((comment) => (
                  <LocketView key={`${locket.name}-${locket.displayTime}`} locket={comment} showCreator />
                ))}
              </>
            )}

            {/* Only show comment editor when user login */}
            {currentUser && (
              <LocketEditor
                key={locket.name}
                cacheKey={`comment-editor-${locket.name}`}
                placeholder={t("editor.add-your-comment-here")}
                parentLocketName={locket.name}
                onConfirm={handleCommentCreated}
              />
            )}
          </div>
        </div>
      </div>
    </section>
  );
};

export default LocketDetail;
