import { Button, Divider, Dropdown, Menu, MenuButton, MenuItem } from "@mui/joy";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import * as api from "@/helpers/api";
import { useTranslate } from "@/utils/i18n";
import showCreateIdentityProviderDialog from "../CreateIdentityProviderDialog";
import { showCommonDialog } from "../Dialog/CommonDialog";
import Icon from "../Icon";
import LearnMore from "../LearnMore";

const SSOSection = () => {
  const t = useTranslate();
  const [identityProviderList, setIdentityProviderList] = useState<IdentityProvider[]>([]);

  useEffect(() => {
    fetchIdentityProviderList();
  }, []);

  const fetchIdentityProviderList = async () => {
    const { data: identityProviderList } = await api.getIdentityProviderList();
    setIdentityProviderList(identityProviderList);
  };

  const handleDeleteIdentityProvider = async (identityProvider: IdentityProvider) => {
    const content = t("setting.sso-section.confirm-delete", { name: identityProvider.name });

    showCommonDialog({
      title: t("setting.sso-section.delete-sso"),
      content: content,
      style: "danger",
      dialogName: "delete-identity-provider-dialog",
      onConfirm: async () => {
        try {
          await api.deleteIdentityProvider(identityProvider.id);
        } catch (error: any) {
          console.error(error);
          toast.error(error.response.data.message);
        }
        await fetchIdentityProviderList();
      },
    });
  };

  return (
    <div className="w-full flex flex-col gap-2 pt-2 pb-4">
      <div className="w-full flex flex-row justify-between items-center gap-1">
        <div className="flex flex-row items-center gap-1">
          <span className="font-mono text-gray-400">{t("setting.sso-section.sso-list")}</span>
          <LearnMore url="https://duyquys.id.vn/docs/advanced-settings/keycloak" />
        </div>
        <Button onClick={() => showCreateIdentityProviderDialog(undefined, fetchIdentityProviderList)}>{t("common.create")}</Button>
      </div>
      <Divider />
      {identityProviderList.map((identityProvider) => (
        <div
          key={identityProvider.id}
          className="py-2 w-full border-b last:border-b dark:border-zinc-700 flex flex-row items-center justify-between"
        >
          <div className="flex flex-row items-center">
            <p className="ml-2">
              {identityProvider.name}
              <span className="text-sm ml-1 opacity-40">({identityProvider.type})</span>
            </p>
          </div>
          <div className="flex flex-row items-center">
            <Dropdown>
              <MenuButton size="sm">
                <Icon.MoreVertical className="w-4 h-auto" />
              </MenuButton>
              <Menu placement="bottom-end" size="sm">
                <MenuItem onClick={() => showCreateIdentityProviderDialog(identityProvider, fetchIdentityProviderList)}>
                  {t("common.edit")}
                </MenuItem>
                <MenuItem onClick={() => handleDeleteIdentityProvider(identityProvider)}>{t("common.delete")}</MenuItem>
              </Menu>
            </Dropdown>
          </div>
        </div>
      ))}
      {identityProviderList.length === 0 && (
        <div className="w-full mt-2 text-sm dark:border-zinc-700 opacity-60 flex flex-row items-center justify-between">
          <p className="">No SSO found.</p>
        </div>
      )}
    </div>
  );
};

export default SSOSection;
