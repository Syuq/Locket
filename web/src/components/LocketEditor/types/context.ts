import { createContext } from "react";
import { LocketRelation } from "@/types/proto/api/v2/locket_relation_service";

interface Context {
  relationList: LocketRelation[];
  setRelationList: (relationList: LocketRelation[]) => void;
  locketName?: string;
}

export const LocketEditorContext = createContext<Context>({
  relationList: [],
  setRelationList: () => {},
});
