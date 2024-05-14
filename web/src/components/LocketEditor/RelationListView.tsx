import { useEffect, useState } from "react";
import { useLocketStore } from "@/store/v1";
import { LocketRelation, LocketRelation_Type } from "@/types/proto/api/v2/locket_relation_service";
import { Locket } from "@/types/proto/api/v2/locket_service";
import Icon from "../Icon";

interface Props {
  relationList: LocketRelation[];
  setRelationList: (relationList: LocketRelation[]) => void;
}

const RelationListView = (props: Props) => {
  const { relationList, setRelationList } = props;
  const locketStore = useLocketStore();
  const [referencingLocketList, setReferencingLocketList] = useState<Locket[]>([]);

  useEffect(() => {
    (async () => {
      const requests = relationList
        .filter((relation) => relation.type === LocketRelation_Type.REFERENCE)
        .map(async (relation) => {
          return await locketStore.getOrFetchLocketByName(relation.relatedLocket, { skipStore: true });
        });
      const list = await Promise.all(requests);
      setReferencingLocketList(list);
    })();
  }, [relationList]);

  const handleDeleteRelation = async (locket: Locket) => {
    setRelationList(relationList.filter((relation) => relation.relatedLocket !== locket.name));
  };

  return (
    <>
      {referencingLocketList.length > 0 && (
        <div className="w-full flex flex-row gap-2 mt-2 flex-wrap">
          {referencingLocketList.map((locket) => {
            return (
              <div
                key={locket.name}
                className="w-auto max-w-xs overflow-hidden flex flex-row justify-start items-center bg-zinc-100 dark:bg-zinc-900 hover:opacity-80 rounded-md text-sm p-1 px-2 text-gray-500 dark:text-gray-400 cursor-pointer hover:line-through"
                onClick={() => handleDeleteRelation(locket)}
              >
                <Icon.Link className="w-4 h-auto shrink-0 opacity-80" />
                <span className="mx-1 max-w-full text-ellipsis whitespace-nowrap overflow-hidden">{locket.content}</span>
                <Icon.X className="w-4 h-auto cursor-pointer shrink-0 opacity-60 hover:opacity-100" />
              </div>
            );
          })}
        </div>
      )}
    </>
  );
};

export default RelationListView;
