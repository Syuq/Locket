import { Dropdown, Menu, MenuButton, MenuItem } from "@mui/joy";
import classNames from "classnames";
import toast from "react-hot-toast";
import { useLocation } from "react-router-dom";
import Icon from "@/components/Icon";
import useNavigateTo from "@/hooks/useNavigateTo";
import { extractLocketIdFromName, useLocketStore } from "@/store/v1";
import { RowStatus } from "@/types/proto/api/v2/common";
import { Locket } from "@/types/proto/api/v2/locket_service";
import { useTranslate } from "@/utils/i18n";
import { showCommonDialog } from "./Dialog/CommonDialog";
import showLocketEditorDialog from "./LocketEditor/LocketEditorDialog";
import showShareLocketDialog from "./ShareLocketDialog";

interface Props {
  locket: Locket;
  className?: string;
  hiddenActions?: ("edit" | "archive" | "delete" | "share" | "pin")[];
}

const LocketActionMenu = (props: Props) => {
  const { locket, hiddenActions } = props;
  const t = useTranslate();
  const location = useLocation();
  const navigateTo = useNavigateTo();
  const locketStore = useLocketStore();
  const isInLocketDetailPage = location.pathname.startsWith(`/m/${locket.uid}`);

  const handleTogglePinLocketBtnClick = async () => {
    try {
      if (locket.pinned) {
        await locketStore.updateLocket(
          {
            name: locket.name,
            pinned: false,
          },
          ["pinned"],
        );
      } else {
        await locketStore.updateLocket(
          {
            name: locket.name,
            pinned: true,
          },
          ["pinned"],
        );
      }
    } catch (error) {
      // do nth
    }
  };

  const handleEditLocketClick = () => {
    showLocketEditorDialog({
      locketName: locket.name,
      cacheKey: `${locket.name}-${locket.displayTime}`,
    });
  };

  const handleToggleLocketStatusClick = async () => {
    try {
      if (locket.rowStatus === RowStatus.ARCHIVED) {
        await locketStore.updateLocket(
          {
            name: locket.name,
            rowStatus: RowStatus.ACTIVE,
          },
          ["row_status"],
        );
        toast(t("message.restored-successfully"));
      } else {
        await locketStore.updateLocket(
          {
            name: locket.name,
            rowStatus: RowStatus.ARCHIVED,
          },
          ["row_status"],
        );
        toast.success(t("message.archived-successfully"));
      }
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
      return;
    }

    if (isInLocketDetailPage) {
      locket.rowStatus === RowStatus.ARCHIVED ? navigateTo("/") : navigateTo("/archived");
    }
  };

  const handleDeleteLocketClick = async () => {
    showCommonDialog({
      title: t("locket.delete-locket"),
      content: t("locket.delete-confirm"),
      style: "danger",
      dialogName: "delete-locket-dialog",
      onConfirm: async () => {
        await locketStore.deleteLocket(locket.name);
        toast.success("Deleted successfully");
        if (isInLocketDetailPage) {
          navigateTo("/");
        }
      },
    });
  };

  return (
    <Dropdown>
      <MenuButton slots={{ root: "div" }}>
        <span className={classNames("flex justify-center items-center rounded-full hover:opacity-70", props.className)}>
          <Icon.MoreVertical className="w-4 h-4 mx-auto text-gray-500 dark:text-gray-400" />
        </span>
      </MenuButton>
      <Menu className="text-sm" size="sm" placement="bottom-end">
        {!hiddenActions?.includes("pin") && (
          <MenuItem onClick={handleTogglePinLocketBtnClick}>
            {locket.pinned ? <Icon.BookmarkMinus className="w-4 h-auto" /> : <Icon.BookmarkPlus className="w-4 h-auto" />}
            {locket.pinned ? t("common.unpin") : t("common.pin")}
          </MenuItem>
        )}
        {!hiddenActions?.includes("edit") && (
          <MenuItem onClick={handleEditLocketClick}>
            <Icon.Edit3 className="w-4 h-auto" />
            {t("common.edit")}
          </MenuItem>
        )}
        {!hiddenActions?.includes("share") && (
          <MenuItem onClick={() => showShareLocketDialog(extractLocketIdFromName(locket.name))}>
            <Icon.Share className="w-4 h-auto" />
            {t("common.share")}
          </MenuItem>
        )}
        <MenuItem color="warning" onClick={handleToggleLocketStatusClick}>
          {locket.rowStatus === RowStatus.ARCHIVED ? (
            <Icon.ArchiveRestore className="w-4 h-auto" />
          ) : (
            <Icon.Archive className="w-4 h-auto" />
          )}
          {locket.rowStatus === RowStatus.ARCHIVED ? t("common.restore") : t("common.archive")}
        </MenuItem>
        <MenuItem color="danger" onClick={handleDeleteLocketClick}>
          <Icon.Trash className="w-4 h-auto" />
          {t("common.delete")}
        </MenuItem>
      </Menu>
    </Dropdown>
  );
};

export default LocketActionMenu;
