import { Tooltip } from "@mui/joy";
import { memo, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useLocketStore } from "@/store/v1";
import { LocketRelation } from "@/types/proto/api/v2/locket_relation_service";
import { Locket } from "@/types/proto/api/v2/locket_service";
import Icon from "./Icon";

interface Props {
  locket: Locket;
  relations: LocketRelation[];
}

const LocketRelationListView = (props: Props) => {
  const { locket, relations: relationList } = props;
  const locketStore = useLocketStore();
  const [referencingLocketList, setReferencingLocketList] = useState<Locket[]>([]);
  const [referencedLocketList, setReferencedLocketList] = useState<Locket[]>([]);

  useEffect(() => {
    (async () => {
      const referencingLocketList = await Promise.all(
        relationList
          .filter((relation) => relation.locket === locket.name && relation.relatedLocket !== locket.name)
          .map((relation) => locketStore.getOrFetchLocketByName(relation.relatedLocket, { skipStore: true })),
      );
      setReferencingLocketList(referencingLocketList);
      const referencedLocketList = await Promise.all(
        relationList
          .filter((relation) => relation.locket !== locket.name && relation.relatedLocket === locket.name)
          .map((relation) => locketStore.getOrFetchLocketByName(relation.locket, { skipStore: true })),
      );
      setReferencedLocketList(referencedLocketList);
    })();
  }, [locket, relationList]);

  return (
    <>
      {referencingLocketList.length > 0 && (
        <div className="w-full flex flex-row justify-start items-center flex-wrap gap-2">
          {referencingLocketList.map((locket) => {
            return (
              <div key={locket.name} className="block w-auto max-w-[50%]">
                <Link
                  className="px-2 border rounded-md w-auto text-sm leading-6 flex flex-row justify-start items-center flex-nowrap text-gray-600 dark:text-gray-400 dark:border-zinc-700 dark:bg-zinc-900 hover:shadow hover:opacity-80"
                  to={`/m/${locket.uid}`}
                  unstable_viewTransition
                >
                  <Tooltip title="Reference" placement="top">
                    <Icon.Link className="w-4 h-auto shrink-0 opacity-70" />
                  </Tooltip>
                  <span className="truncate ml-1">{locket.content}</span>
                </Link>
              </div>
            );
          })}
        </div>
      )}
      {referencedLocketList.length > 0 && (
        <div className="w-full flex flex-row justify-start items-center flex-wrap gap-2">
          {referencedLocketList.map((locket) => {
            return (
              <div key={locket.name} className="block w-auto max-w-[50%]">
                <Link
                  className="px-2 border rounded-md w-auto text-sm leading-6 flex flex-row justify-start items-center flex-nowrap text-gray-600 dark:text-gray-400 dark:border-zinc-700 dark:bg-zinc-900 hover:shadow hover:opacity-80"
                  to={`/m/${locket.uid}`}
                  unstable_viewTransition
                >
                  <Tooltip title="Backlink" placement="top">
                    <Icon.Milestone className="w-4 h-auto shrink-0 opacity-70" />
                  </Tooltip>
                  <span className="truncate ml-1">{locket.content}</span>
                </Link>
              </div>
            );
          })}
        </div>
      )}
    </>
  );
};

export default memo(LocketRelationListView);
