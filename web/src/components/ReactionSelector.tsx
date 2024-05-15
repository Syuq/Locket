import { Dropdown, Menu, MenuButton } from "@mui/joy";
import classNames from "classnames";
import { useRef, useState } from "react";
import useClickAway from "react-use/lib/useClickAway";
import Icon from "@/components/Icon";
import { locketServiceClient } from "@/grpcweb";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useLocketStore } from "@/store/v1";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { Reaction_Type } from "@/types/proto/api/v2/reaction_service";
import { stringifyReactionType } from "./ReactionView";

interface Props {
  locket: Locket;
  className?: string;
}

const REACTION_TYPES = [
  Reaction_Type.THUMBS_UP,
  Reaction_Type.THUMBS_DOWN,
  Reaction_Type.HEART,
  Reaction_Type.FIRE,
  Reaction_Type.CLAPPING_HANDS,
  Reaction_Type.LAUGH,
  Reaction_Type.OK_HAND,
  Reaction_Type.ROCKET,
  Reaction_Type.EYES,
  Reaction_Type.THINKING_FACE,
  Reaction_Type.CLOWN_FACE,
  Reaction_Type.QUESTION_MARK,
];

const ReactionSelector = (props: Props) => {
  const { locket, className } = props;
  const currentUser = useCurrentUser();
  const locketStore = useLocketStore();
  const [open, setOpen] = useState(false);
  const containerRef = useRef<HTMLDivElement>(null);

  useClickAway(containerRef, () => {
    setOpen(false);
  });

  const hasReacted = (reactionType: Reaction_Type) => {
    return locket.reactions.some((r) => r.reactionType === reactionType && r.creator === currentUser?.name);
  };

  const handleReactionClick = async (reactionType: Reaction_Type) => {
    try {
      if (hasReacted(reactionType)) {
        const reactions = locket.reactions.filter(
          (reaction) => reaction.reactionType === reactionType && reaction.creator === currentUser.name,
        );
        for (const reaction of reactions) {
          await locketServiceClient.deleteLocketReaction({ reactionId: reaction.id });
        }
      } else {
        await locketServiceClient.upsertLocketReaction({
          name: locket.name,
          reaction: {
            contentId: locket.name,
            reactionType: reactionType,
          },
        });
      }
      await locketStore.getOrFetchLocketByName(locket.name, { skipCache: true });
    } catch (error) {
      // skip error.
    }
    setOpen(false);
  };

  return (
    <Dropdown open={open} onOpenChange={(_, isOpen) => setOpen(isOpen)}>
      <MenuButton slots={{ root: "div" }}>
        <span
          className={classNames(
            "h-7 w-7 flex justify-center items-center rounded-full border dark:border-zinc-700 hover:opacity-70",
            className,
          )}
        >
          <Icon.SmilePlus className="w-4 h-4 mx-auto text-gray-500 dark:text-gray-400" />
        </span>
      </MenuButton>
      <Menu className="relative text-sm" component="div" size="sm" placement="bottom-start">
        <div ref={containerRef}>
          <div className="grid grid-cols-6 py-0.5 px-2 h-auto font-mono gap-1">
            {REACTION_TYPES.map((reactionType) => {
              return (
                <span
                  key={reactionType}
                  className={classNames(
                    "inline-flex w-auto cursor-pointer rounded text-lg px-1 text-gray-500 dark:text-gray-400 hover:opacity-80",
                    hasReacted(reactionType) && "bg-blue-100 dark:bg-zinc-800",
                  )}
                  onClick={() => handleReactionClick(reactionType)}
                >
                  {stringifyReactionType(reactionType)}
                </span>
              );
            })}
          </div>
        </div>
      </Menu>
    </Dropdown>
  );
};

export default ReactionSelector;
