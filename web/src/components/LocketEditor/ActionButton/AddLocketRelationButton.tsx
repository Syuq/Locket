import { IconButton } from "@mui/joy";
import { uniqBy } from "lodash-es";
import { useContext } from "react";
import toast from "react-hot-toast";
import showCreateLocketRelationDialog from "@/components/CreateLocketRelationDialog";
import Icon from "@/components/Icon";
import { LocketRelation_Type } from "@/types/proto/api/v2/locket_relation_service";
import { EditorRefActions } from "../Editor";
import { LocketEditorContext } from "../types";

interface Props {
  editorRef: React.RefObject<EditorRefActions>;
}

const AddLocketRelationButton = (props: Props) => {
  const { editorRef } = props;
  const context = useContext(LocketEditorContext);

  const handleAddLocketRelationBtnClick = () => {
    showCreateLocketRelationDialog({
      onConfirm: (lockets, embedded) => {
        // If embedded mode is enabled, embed the locket instead of creating a relation.
        if (embedded) {
          if (!editorRef.current) {
            toast.error("Failed to embed locket");
            return;
          }

          const cursorPosition = editorRef.current.getCursorPosition();
          const prevValue = editorRef.current.getContent().slice(0, cursorPosition);
          if (prevValue !== "" && !prevValue.endsWith("\n")) {
            editorRef.current.insertText("\n");
          }
          for (const locket of lockets) {
            editorRef.current.insertText(`![[lockets/${locket.uid}]]\n`);
          }
          setTimeout(() => {
            editorRef.current?.scrollToCursor();
            editorRef.current?.focus();
          });
          return;
        }

        context.setRelationList(
          uniqBy(
            [
              ...lockets.map((locket) => ({
                locket: context.locketName || "",
                relatedLocket: locket.name,
                type: LocketRelation_Type.REFERENCE,
              })),
              ...context.relationList,
            ].filter((relation) => relation.relatedLocket !== context.locketName),
            "relatedLocketId",
          ),
        );
      },
    });
  };

  return (
    <IconButton size="sm" onClick={handleAddLocketRelationBtnClick}>
      <Icon.Link className="w-5 h-5 mx-auto" />
    </IconButton>
  );
};

export default AddLocketRelationButton;
