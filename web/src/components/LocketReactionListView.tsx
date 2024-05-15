import { uniq } from "lodash-es";
import { memo, useEffect, useState } from "react";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useUserStore } from "@/store/v1";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { Reaction, Reaction_Type } from "@/types/proto/api/v2/reaction_service";
import { User } from "@/types/proto/api/v2/user_service";
import ReactionSelector from "./ReactionSelector";
import ReactionView from "./ReactionView";

interface Props {
  locket: Locket;
  reactions: Reaction[];
}

const LocketReactionListView = (props: Props) => {
  const { locket, reactions } = props;
  const currentUser = useCurrentUser();
  const userStore = useUserStore();
  const [reactionGroup, setReactionGroup] = useState<Map<Reaction_Type, User[]>>(new Map());

  useEffect(() => {
    (async () => {
      const reactionGroup = new Map<Reaction_Type, User[]>();
      for (const reaction of reactions) {
        const user = await userStore.getOrFetchUserByName(reaction.creator);
        const users = reactionGroup.get(reaction.reactionType) || [];
        users.push(user);
        reactionGroup.set(reaction.reactionType, uniq(users));
      }
      setReactionGroup(reactionGroup);
    })();
  }, [reactions]);

  return (
    reactions.length > 0 && (
      <div className="w-full flex flex-row justify-start items-start flex-wrap gap-1 select-none">
        {Array.from(reactionGroup).map(([reactionType, users]) => {
          return (
            <ReactionView key={`${reactionType.toString()} ${users.length}`} locket={locket} reactionType={reactionType} users={users} />
          );
        })}
        {currentUser && <ReactionSelector locket={locket} />}
      </div>
    )
  );
};

export default memo(LocketReactionListView);
