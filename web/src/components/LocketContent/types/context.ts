import { createContext } from "react";
import { Node } from "@/types/node";

interface Context {
  nodes: Node[];
  // embeddedLockets is a set of locket resource names that are embedded in the current locket.
  // This is used to prevent infinite loops when a locket embeds itself.
  embeddedLockets: Set<string>;
  locketName?: string;
  readonly?: boolean;
  disableFilter?: boolean;
}

export const RendererContext = createContext<Context>({
  nodes: [],
  embeddedLockets: new Set(),
});
