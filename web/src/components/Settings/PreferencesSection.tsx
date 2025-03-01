import { Button, Divider, Input, Option, Select } from "@mui/joy";
import { useState } from "react";
import { toast } from "react-hot-toast";
import { useGlobalStore } from "@/store/module";
import { useUserStore } from "@/store/v1";
import { Visibility } from "@/types/proto/api/v2/locket_service";
import { UserSetting } from "@/types/proto/api/v2/user_service";
import { useTranslate } from "@/utils/i18n";
import { convertVisibilityFromString, convertVisibilityToString } from "@/utils/locket";
import AppearanceSelect from "../AppearanceSelect";
import LocaleSelect from "../LocaleSelect";
import VisibilityIcon from "../VisibilityIcon";
import WebhookSection from "./WebhookSection";

const PreferencesSection = () => {
  const t = useTranslate();
  const globalStore = useGlobalStore();
  const userStore = useUserStore();
  const setting = userStore.userSetting as UserSetting;
  const [telegramUserId, setTelegramUserId] = useState<string>(setting.telegramUserId);

  const handleLocaleSelectChange = async (locale: Locale) => {
    await userStore.updateUserSetting(
      {
        locale,
      },
      ["locale"],
    );
    globalStore.setLocale(locale);
  };

  const handleAppearanceSelectChange = async (appearance: Appearance) => {
    await userStore.updateUserSetting(
      {
        appearance,
      },
      ["appearance"],
    );
    globalStore.setAppearance(appearance);
  };

  const handleDefaultLocketVisibilityChanged = async (value: string) => {
    await userStore.updateUserSetting(
      {
        locketVisibility: value,
      },
      ["locket_visibility"],
    );
  };

  const handleSaveTelegramUserId = async () => {
    try {
      await userStore.updateUserSetting(
        {
          telegramUserId: telegramUserId,
        },
        ["telegram_user_id"],
      );
      toast.success(t("message.update-succeed"));
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
    }
  };

  const handleTelegramUserIdChanged = async (value: string) => {
    setTelegramUserId(value);
  };

  return (
    <div className="w-full flex flex-col gap-2 pt-2 pb-4">
      <p className="font-medium text-gray-700 dark:text-gray-500">{t("common.basic")}</p>
      <div className="w-full flex flex-row justify-between items-center">
        <span>{t("common.language")}</span>
        <LocaleSelect value={setting.locale} onChange={handleLocaleSelectChange} />
      </div>
      <div className="w-full flex flex-row justify-between items-center">
        <span>{t("setting.preference-section.theme")}</span>
        <AppearanceSelect value={setting.appearance as Appearance} onChange={handleAppearanceSelectChange} />
      </div>
      <p className="font-medium text-gray-700 dark:text-gray-500">{t("setting.preference")}</p>
      <div className="w-full flex flex-row justify-between items-center">
        <span className="truncate">{t("setting.preference-section.default-locket-visibility")}</span>
        <Select
          className="!min-w-fit"
          value={setting.locketVisibility}
          startDecorator={<VisibilityIcon visibility={convertVisibilityFromString(setting.locketVisibility)} />}
          onChange={(_, visibility) => {
            if (visibility) {
              handleDefaultLocketVisibilityChanged(visibility);
            }
          }}
        >
          {[Visibility.PRIVATE, Visibility.PROTECTED, Visibility.PUBLIC]
            .map((v) => convertVisibilityToString(v))
            .map((item) => (
              <Option key={item} value={item} className="whitespace-nowrap">
                {t(`locket.visibility.${item.toLowerCase() as Lowercase<typeof item>}`)}
              </Option>
            ))}
        </Select>
      </div>

      <Divider className="!my-3" />

      <div className="space-y-2 border rounded-md py-2 px-3 dark:border-zinc-700">
        <div className="w-full flex flex-row justify-between items-center">
          <div className="w-auto flex items-center">
            <span className="mr-1">{t("setting.preference-section.telegram-user-id")}</span>
          </div>
          <Button variant="outlined" color="neutral" onClick={handleSaveTelegramUserId}>
            {t("common.save")}
          </Button>
        </div>
        <Input
          className="w-full"
          sx={{
            fontFamily: "monospace",
            fontSize: "14px",
          }}
          value={telegramUserId}
          onChange={(event) => handleTelegramUserIdChanged(event.target.value)}
          placeholder={t("setting.preference-section.telegram-user-id-placeholder")}
        />
      </div>

      <Divider className="!my-3" />

      <WebhookSection />
    </div>
  );
};

export default PreferencesSection;
