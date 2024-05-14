import { useEffect } from "react";
import useLoading from "@/hooks/useLoading";
import useNavigateTo from "@/hooks/useNavigateTo";
import { useLocketStore } from "@/store/v1";
import Error from "./Error";

interface Props {
  resourceId: string;
  params: string;
}

const ReferencedLocket = ({ resourceId, params: paramsStr }: Props) => {
  const navigateTo = useNavigateTo();
  const loadingState = useLoading();
  const locketStore = useLocketStore();
  const locket = locketStore.getLocketByUid(resourceId);
  const params = new URLSearchParams(paramsStr);

  useEffect(() => {
    locketStore.searchLockets(`uid == "${resourceId}"`).finally(() => loadingState.setFinish());
  }, [resourceId]);

  if (loadingState.isLoading) {
    return null;
  }
  if (!locket) {
    return <Error message={`Locket not found: ${resourceId}`} />;
  }

  const paramsText = params.has("text") ? params.get("text") : undefined;
  const displayContent = paramsText || (locket.content.length > 12 ? `${locket.content.slice(0, 12)}...` : locket.content);

  const handleGotoLocketDetailPage = () => {
    navigateTo(`/m/${locket.uid}`);
  };

  return (
    <span
      className="text-blue-600 whitespace-nowrap dark:text-blue-400 cursor-pointer underline break-all hover:opacity-80 decoration-1"
      onClick={handleGotoLocketDetailPage}
    >
      {displayContent}
    </span>
  );
};

export default ReferencedLocket;
