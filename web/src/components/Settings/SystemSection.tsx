import { Button, Divider, Input, Switch, Textarea, Tooltip } from "@mui/joy";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { workspaceSettingServiceClient } from "@/grpcweb";
import * as api from "@/helpers/api";
import { useGlobalStore } from "@/store/module";
import { WorkspaceSettingPrefix } from "@/store/v1";
import { WorkspaceGeneralSetting } from "@/types/proto/api/v2/workspace_setting_service";
import { WorkspaceSettingKey } from "@/types/proto/store/workspace_setting";
import { useTranslate } from "@/utils/i18n";
import { showCommonDialog } from "../Dialog/CommonDialog";
import Icon from "../Icon";
import showUpdateCustomizedProfileDialog from "../UpdateCustomizedProfileDialog";

interface State {
  disablePublicLockets: boolean;
  maxUploadSizeMiB: number;
  locketDisplayWithUpdatedTs: boolean;
}

const SystemSection = () => {
  const t = useTranslate();
  const globalStore = useGlobalStore();
  const systemStatus = globalStore.state.systemStatus;
  const [state, setState] = useState<State>({
    disablePublicLockets: systemStatus.disablePublicLockets,
    maxUploadSizeMiB: systemStatus.maxUploadSizeMiB,
    locketDisplayWithUpdatedTs: systemStatus.locketDisplayWithUpdatedTs,
  });
  const [workspaceGeneralSetting, setWorkspaceGeneralSetting] = useState<WorkspaceGeneralSetting>(WorkspaceGeneralSetting.fromPartial({}));
  const [telegramBotToken, setTelegramBotToken] = useState<string>("");

  useEffect(() => {
    (async () => {
      await globalStore.fetchSystemStatus();
      const { setting } = await workspaceSettingServiceClient.getWorkspaceSetting({
        name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
      });
      if (setting && setting.generalSetting) {
        setWorkspaceGeneralSetting(WorkspaceGeneralSetting.fromPartial(setting.generalSetting));
      }
    })();
  }, []);

  useEffect(() => {
    api.getSystemSetting().then(({ data: systemSettings }) => {
      const telegramBotSetting = systemSettings.find((setting) => setting.name === "telegram-bot-token");
      if (telegramBotSetting) {
        setTelegramBotToken(telegramBotSetting.value);
      }
    });
  }, []);

  useEffect(() => {
    setState({
      ...state,
      disablePublicLockets: systemStatus.disablePublicLockets,
      maxUploadSizeMiB: systemStatus.maxUploadSizeMiB,
      locketDisplayWithUpdatedTs: systemStatus.locketDisplayWithUpdatedTs,
    });
  }, [systemStatus]);

  const handleAllowSignUpChanged = async (value: boolean) => {
    const setting = { ...workspaceGeneralSetting, disallowSignup: !value };
    await workspaceSettingServiceClient.setWorkspaceSetting({
      setting: {
        name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
        generalSetting: setting,
      },
    });
    setWorkspaceGeneralSetting(setting);
  };

  const handleDisablePasswordLoginChanged = async (value: boolean) => {
    const updateSetting = async () => {
      const setting = { ...workspaceGeneralSetting, disallowPasswordLogin: value };
      await workspaceSettingServiceClient.setWorkspaceSetting({
        setting: {
          name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
          generalSetting: setting,
        },
      });
      setWorkspaceGeneralSetting(setting);
    };
    if (value) {
      showCommonDialog({
        title: "Confirm",
        content: "Are you sure to disable password login?",
        style: "danger",
        dialogName: "disable-password-login-dialog",
        onConfirm: async () => {
          await updateSetting();
        },
      });
    } else {
      await updateSetting();
    }
  };

  const handleUpdateCustomizedProfileButtonClick = () => {
    showUpdateCustomizedProfileDialog();
  };

  const handleInstanceUrlChanged = (value: string) => {
    setWorkspaceGeneralSetting({ ...workspaceGeneralSetting, instanceUrl: value });
  };

  const handleSaveInstanceUrl = async () => {
    try {
      await workspaceSettingServiceClient.setWorkspaceSetting({
        setting: {
          name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
          generalSetting: workspaceGeneralSetting,
        },
      });
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
      return;
    }
    toast.success("Instance URL updated");
  };

  const handleTelegramBotTokenChanged = (value: string) => {
    setTelegramBotToken(value);
  };

  const handleSaveTelegramBotToken = async () => {
    try {
      await api.upsertSystemSetting({
        name: "telegram-bot-token",
        value: telegramBotToken,
      });
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
      return;
    }
    toast.success("Telegram Bot Token updated");
  };

  const handleAdditionalStyleChanged = (value: string) => {
    setWorkspaceGeneralSetting({ ...workspaceGeneralSetting, additionalStyle: value });
  };

  const handleSaveAdditionalStyle = async () => {
    try {
      await workspaceSettingServiceClient.setWorkspaceSetting({
        setting: {
          name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
          generalSetting: workspaceGeneralSetting,
        },
      });
    } catch (error: any) {
      toast.error(error.response.data.message);
      console.error(error);
      return;
    }
    toast.success(t("message.succeed-update-additional-style"));
  };

  const handleAdditionalScriptChanged = (value: string) => {
    setWorkspaceGeneralSetting({ ...workspaceGeneralSetting, additionalScript: value });
  };

  const handleSaveAdditionalScript = async () => {
    try {
      await workspaceSettingServiceClient.setWorkspaceSetting({
        setting: {
          name: `${WorkspaceSettingPrefix}${WorkspaceSettingKey.WORKSPACE_SETTING_GENERAL}`,
          generalSetting: workspaceGeneralSetting,
        },
      });
    } catch (error: any) {
      toast.error(error.response.data.message);
      console.error(error);
      return;
    }
    toast.success(t("message.succeed-update-additional-script"));
  };

  const handleDisablePublicLocketsChanged = async (value: boolean) => {
    setState({
      ...state,
      disablePublicLockets: value,
    });
    globalStore.setSystemStatus({ disablePublicLockets: value });
    await api.upsertSystemSetting({
      name: "disable-public-lockets",
      value: JSON.stringify(value),
    });
  };

  const handleLocketDisplayWithUpdatedTs = async (value: boolean) => {
    setState({
      ...state,
      locketDisplayWithUpdatedTs: value,
    });
    globalStore.setSystemStatus({ locketDisplayWithUpdatedTs: value });
    await api.upsertSystemSetting({
      name: "locket-display-with-updated-ts",
      value: JSON.stringify(value),
    });
  };

  const handleMaxUploadSizeChanged = async (event: React.FocusEvent<HTMLInputElement>) => {
    // fixes cursor skipping position on mobile
    event.target.selectionEnd = event.target.value.length;

    let num = parseInt(event.target.value);
    if (Number.isNaN(num)) {
      num = 0;
    }
    setState({
      ...state,
      maxUploadSizeMiB: num,
    });
    event.target.value = num.toString();
    globalStore.setSystemStatus({ maxUploadSizeMiB: num });
    await api.upsertSystemSetting({
      name: "max-upload-size-mib",
      value: JSON.stringify(num),
    });
  };

  const handleMaxUploadSizeFocus = (event: React.FocusEvent<HTMLInputElement>) => {
    event.target.select();
  };

  return (
    <div className="w-full flex flex-col gap-2 pt-2 pb-4">
      <p className="font-medium text-gray-700 dark:text-gray-500">{t("common.basic")}</p>
      <div className="w-full flex flex-row justify-between items-center">
        <div>
          {t("setting.system-section.server-name")}: <span className="font-mono font-bold">{systemStatus.customizedProfile.name}</span>
        </div>
        <Button onClick={handleUpdateCustomizedProfileButtonClick}>{t("common.edit")}</Button>
      </div>
      <p className="font-medium text-gray-700 dark:text-gray-500">General</p>
      <div className="w-full flex flex-row justify-between items-center">
        <span className="mr-1">{t("setting.system-section.allow-user-signup")}</span>
        <Switch checked={!workspaceGeneralSetting.disallowSignup} onChange={(event) => handleAllowSignUpChanged(event.target.checked)} />
      </div>
      <div className="w-full flex flex-row justify-between items-center">
        <span className="mr-1">{t("setting.system-section.disable-password-login")}</span>
        <Switch
          checked={workspaceGeneralSetting.disallowPasswordLogin}
          onChange={(event) => handleDisablePasswordLoginChanged(event.target.checked)}
        />
      </div>
      <div className="space-y-2 border rounded-md py-2 px-3 dark:border-zinc-700">
        <div className="w-full flex flex-row justify-between items-center">
          <div className="flex flex-row items-center">
            <div className="w-auto flex items-center">
              <span className="mr-1">Instance URL</span>
            </div>
          </div>
          <Button variant="outlined" color="neutral" onClick={handleSaveInstanceUrl}>
            {t("common.save")}
          </Button>
        </div>
        <Input
          className="w-full"
          sx={{
            fontFamily: "monospace",
            fontSize: "14px",
          }}
          placeholder={"Should be started with http:// or https://"}
          value={workspaceGeneralSetting.instanceUrl}
          onChange={(event) => handleInstanceUrlChanged(event.target.value)}
        />
      </div>
      <div className="space-y-2 border rounded-md py-2 px-3 dark:border-zinc-700">
        <div className="w-full flex flex-row justify-between items-center">
          <span>{t("setting.system-section.additional-style")}</span>
          <Button variant="outlined" color="neutral" onClick={handleSaveAdditionalStyle}>
            {t("common.save")}
          </Button>
        </div>
        <Textarea
          className="w-full"
          sx={{
            fontFamily: "monospace",
            fontSize: "14px",
          }}
          minRows={2}
          maxRows={4}
          placeholder={t("setting.system-section.additional-style-placeholder")}
          value={workspaceGeneralSetting.additionalStyle}
          onChange={(event) => handleAdditionalStyleChanged(event.target.value)}
        />
        <div className="w-full flex flex-row justify-between items-center">
          <span>{t("setting.system-section.additional-script")}</span>
          <Button variant="outlined" color="neutral" onClick={handleSaveAdditionalScript}>
            {t("common.save")}
          </Button>
        </div>
        <Textarea
          className="w-full"
          color="neutral"
          sx={{
            fontFamily: "monospace",
            fontSize: "14px",
          }}
          minRows={2}
          maxRows={4}
          placeholder={t("setting.system-section.additional-script-placeholder")}
          value={workspaceGeneralSetting.additionalScript}
          onChange={(event) => handleAdditionalScriptChanged(event.target.value)}
        />
      </div>
      <Divider className="!my-3" />
      <p className="font-medium text-gray-700 dark:text-gray-500">Others</p>
      <div className="w-full flex flex-row justify-between items-center">
        <span>{t("setting.system-section.disable-public-lockets")}</span>
        <Switch checked={state.disablePublicLockets} onChange={(event) => handleDisablePublicLocketsChanged(event.target.checked)} />
      </div>
      <div className="w-full flex flex-row justify-between items-center">
        <span>{t("setting.system-section.display-with-updated-time")}</span>
        <Switch checked={state.locketDisplayWithUpdatedTs} onChange={(event) => handleLocketDisplayWithUpdatedTs(event.target.checked)} />
      </div>
      <div className="w-full flex flex-row justify-between items-center">
        <div className="flex flex-row items-center">
          <span className="mr-1">{t("setting.system-section.max-upload-size")}</span>
          <Tooltip title={t("setting.system-section.max-upload-size-hint")} placement="top">
            <Icon.HelpCircle className="w-4 h-auto" />
          </Tooltip>
        </div>
        <Input
          className="w-16"
          sx={{
            fontFamily: "monospace",
          }}
          defaultValue={state.maxUploadSizeMiB}
          onFocus={handleMaxUploadSizeFocus}
          onChange={handleMaxUploadSizeChanged}
        />
      </div>
      <div className="space-y-2 border rounded-md py-2 px-3 dark:border-zinc-700">
        <div className="w-full flex flex-row justify-between items-center">
          <div className="flex flex-row items-center">
            <div className="w-auto flex items-center">
              <span className="mr-1">{t("setting.system-section.telegram-bot-token")}</span>
            </div>
          </div>
          <Button variant="outlined" color="neutral" onClick={handleSaveTelegramBotToken}>
            {t("common.save")}
          </Button>
        </div>
        <Input
          className="w-full"
          sx={{
            fontFamily: "monospace",
            fontSize: "14px",
          }}
          placeholder={t("setting.system-section.telegram-bot-token-placeholder")}
          value={telegramBotToken}
          onChange={(event) => handleTelegramBotTokenChanged(event.target.value)}
        />
      </div>
    </div>
  );
};

export default SystemSection;
