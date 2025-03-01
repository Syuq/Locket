import classNames from "classnames";
import { useEffect } from "react";
import LocketResourceListView from "@/components/LocketResourceListView";
import useLoading from "@/hooks/useLoading";
import { useResourceStore } from "@/store/v1";
import Error from "./Error";

interface Props {
  resourceId: string;
  params: string;
}

const getAdditionalClassNameWithParams = (params: URLSearchParams) => {
  const additionalClassNames = [];
  if (params.has("align")) {
    const align = params.get("align");
    if (align === "center") {
      additionalClassNames.push("mx-auto");
    }
  }
  if (params.has("size")) {
    const size = params.get("size");
    if (size === "lg") {
      additionalClassNames.push("w-full");
    } else if (size === "md") {
      additionalClassNames.push("w-2/3");
    } else if (size === "sm") {
      additionalClassNames.push("w-1/3");
    }
  }
  if (params.has("width")) {
    const width = params.get("width");
    additionalClassNames.push(`w-[${width}]`);
  }
  return additionalClassNames.join(" ");
};

const EmbeddedResource = ({ resourceId, params: paramsStr }: Props) => {
  const loadingState = useLoading();
  const resourceStore = useResourceStore();
  const resource = resourceStore.getResourceByName(resourceId);
  const params = new URLSearchParams(paramsStr);

  useEffect(() => {
    resourceStore.searchResources(`uid == ${resourceId}`).finally(() => loadingState.setFinish());
  }, [resourceId]);

  if (loadingState.isLoading) {
    return null;
  }
  if (!resource) {
    return <Error message={`Resource not found: ${resourceId}`} />;
  }

  return (
    <div className={classNames("max-w-full", getAdditionalClassNameWithParams(params))}>
      <LocketResourceListView resources={[resource]} />
    </div>
  );
};

export default EmbeddedResource;
