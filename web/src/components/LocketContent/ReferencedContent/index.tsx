import Error from "./Error";
import ReferencedLocket from "./ReferencedLocket";

interface Props {
  resourceName: string;
  params: string;
}

const extractResourceTypeAndId = (resourceName: string) => {
  const [resourceType, resourceId] = resourceName.split("/");
  return { resourceType, resourceId };
};

const ReferencedContent = ({ resourceName, params }: Props) => {
  const { resourceType, resourceId } = extractResourceTypeAndId(resourceName);
  if (resourceType === "lockets") {
    return <ReferencedLocket resourceId={resourceId} params={params} />;
  }
  return <Error message={`Unknown resource: ${resourceName}`} />;
};

export default ReferencedContent;
