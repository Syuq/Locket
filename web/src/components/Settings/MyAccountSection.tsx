import { Button, Dropdown, Menu, MenuButton, MenuItem } from "@mui/joy";
import { locketServiceClient } from "@/grpcweb";
import { downloadFileFromUrl } from "@/helpers/utils";
import useCurrentUser from "@/hooks/useCurrentUser";
import { useTranslate } from "@/utils/i18n";
import showChangePasswordDialog from "../ChangePasswordDialog";
import Icon from "../Icon";
import showUpdateAccountDialog from "../UpdateAccountDialog";
import UserAvatar from "../UserAvatar";
import AccessTokenSection from "./AccessTokenSection";

const MyAccountSection = () => {
  const t = useTranslate();
  const user = useCurrentUser();

  const downloadExportedLockets = async (user: any) => {
    const { content } = await locketServiceClient.exportLockets({ filter: `creator == "${user.name}"` });
    const downloadUrl = window.URL.createObjectURL(new Blob([content]));
    downloadFileFromUrl(downloadUrl, "lockets-export.zip");
    URL.revokeObjectURL(downloadUrl);
  };

  return (
    <div className="w-full gap-2 pt-2 pb-4">
      <p className="font-medium text-gray-700 dark:text-gray-500">{t("setting.account-section.title")}</p>
      <div className="w-full mt-2 flex flex-row justify-start items-center">
        <UserAvatar className="mr-2 shrink-0 w-10 h-10" avatarUrl={user.avatarUrl} />
        <div className="max-w-[calc(100%-3rem)] flex flex-col justify-center items-start">
          <p className="w-full">
            <span className="text-xl leading-tight font-medium">{user.nickname}</span>
            <span className="ml-1 text-base leading-tight text-gray-500 dark:text-gray-400">({user.username})</span>
          </p>
          <p className="w-4/5 leading-tight text-sm truncate">{user.description}</p>
        </div>
      </div>
      <div className="w-full flex flex-row justify-start items-center mt-2 space-x-2">
        <Button variant="outlined" color="neutral" size="sm" onClick={showUpdateAccountDialog}>
          <Icon.PenLine className="w-4 h-4 mx-auto mr-1" />
          {t("common.edit")}
        </Button>
        <Dropdown>
          <MenuButton slots={{ root: "div" }}>
            <Button variant="outlined" color="neutral" size="sm">
              <Icon.MoreVertical className="w-4 h-4 mx-auto" />
            </Button>
          </MenuButton>
          <Menu className="text-sm" size="sm" placement="bottom">
            <MenuItem onClick={showChangePasswordDialog}>{t("setting.account-section.change-password")}</MenuItem>
            <MenuItem onClick={() => downloadExportedLockets(user)}>{t("setting.account-section.export-lockets")}</MenuItem>
          </Menu>
        </Dropdown>
      </div>

      <AccessTokenSection />
    </div>
  );
};

export default MyAccountSection;
