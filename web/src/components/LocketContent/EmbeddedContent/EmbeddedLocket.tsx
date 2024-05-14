import { useContext, useEffect } from "react";
import { Link } from "react-router-dom";
import Icon from "@/components/Icon";
import LocketResourceListView from "@/components/LocketResourceListView";
import useLoading from "@/hooks/useLoading";
import { useLocketStore } from "@/store/v1";
import LocketContent from "..";
import { RendererContext } from "../types";
import Error from "./Error";

interface Props {
  resourceId: string;
  params: string;
}

const EmbeddedLocket = ({ resourceId, params: paramsStr }: Props) => {
  const context = useContext(RendererContext);
  const loadingState = useLoading();
  const locketStore = useLocketStore();
  const locket = locketStore.getLocketByUid(resourceId);
  const resourceName = `lockets/${resourceId}`;

  useEffect(() => {
    locketStore.searchLockets(`uid == "${resourceId}"`).finally(() => loadingState.setFinish());
  }, [resourceId]);

  if (loadingState.isLoading) {
    return null;
  }
  if (!locket) {
    return <Error message={`Locket not found: ${resourceId}`} />;
  }
  if (locket.name === context.locketName || context.embeddedLockets.has(resourceName)) {
    return <Error message={`Nested Rendering Error: ![[${resourceName}]]`} />;
  }

  // Add the locket to the set of embedded lockets. This is used to prevent infinite loops when a locket embeds itself.
  context.embeddedLockets.add(resourceName);
  const params = new URLSearchParams(paramsStr);
  const inlineMode = params.has("inline");
  if (inlineMode) {
    return (
      <div className="w-full">
        <LocketContent
          key={`${locket.name}-${locket.updateTime}`}
          locketName={locket.name}
          content={locket.content}
          embeddedLockets={context.embeddedLockets}
        />
        <LocketResourceListView resources={locket.resources} />
      </div>
    );
  }

  return (
    <div className="relative flex flex-col justify-start items-start w-full px-3 py-2 bg-white dark:bg-zinc-800 rounded-lg border border-gray-200 dark:border-zinc-700 hover:shadow">
      <div className="w-full mb-1 flex flex-row justify-between items-center">
        <div className="text-sm leading-6 text-gray-400 select-none">
          <relative-time datetime={locket.displayTime?.toISOString()} format="datetime" tense="past"></relative-time>
        </div>
        <Link className="hover:opacity-80" to={`/m/${locket.uid}`} unstable_viewTransition>
          <Icon.ArrowUpRight className="w-5 h-auto opacity-80 text-gray-400" />
        </Link>
      </div>
      <LocketContent
        key={`${locket.name}-${locket.updateTime}`}
        locketName={locket.name}
        content={locket.content}
        embeddedLockets={context.embeddedLockets}
      />
      <LocketResourceListView resources={locket.resources} />
    </div>
  );
};

export default EmbeddedLocket;
