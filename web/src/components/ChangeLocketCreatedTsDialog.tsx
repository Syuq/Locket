import { Button, IconButton, Input } from "@mui/joy";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { getNormalizedTimeString } from "@/helpers/datetime";
import { LocketNamePrefix, useLocketStore } from "@/store/v1";
import { useTranslate } from "@/utils/i18n";
import { generateDialog } from "./Dialog";
import Icon from "./Icon";

interface Props extends DialogProps {
  locketId: number;
}

const ChangeLocketCreatedTsDialog: React.FC<Props> = (props: Props) => {
  const t = useTranslate();
  const { destroy, locketId } = props;
  const locketStore = useLocketStore();
  const [createdAt, setCreatedAt] = useState("");
  const maxDatetimeValue = getNormalizedTimeString();

  useEffect(() => {
    locketStore.getOrFetchLocketByName(`${LocketNamePrefix}${locketId}`).then((locket) => {
      if (locket) {
        const datetime = getNormalizedTimeString(locket.createTime);
        setCreatedAt(datetime);
      } else {
        toast.error(t("message.locket-not-found"));
        destroy();
      }
    });
  }, []);

  const handleCloseBtnClick = () => {
    destroy();
  };

  const handleDatetimeInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const datetime = e.target.value as string;
    setCreatedAt(datetime);
  };

  const handleSaveBtnClick = async () => {
    try {
      await locketStore.updateLocket(
        {
          name: `${LocketNamePrefix}${locketId}`,
          createTime: new Date(createdAt),
        },
        ["created_ts"],
      );
      toast.success(t("message.locket-updated-datetime"));
      handleCloseBtnClick();
    } catch (error: any) {
      console.error(error);
      toast.error(error.response.data.message);
    }
  };

  return (
    <>
      <div className="dialog-header-container">
        <p className="title-text">{t("message.change-locket-created-time")}</p>
        <IconButton size="sm" onClick={handleCloseBtnClick}>
          <Icon.X className="w-5 h-auto" />
        </IconButton>
      </div>
      <div className="flex flex-col justify-start items-start !w-72 max-w-full">
        <Input
          className="w-full"
          type="datetime-local"
          value={createdAt}
          slotProps={{
            input: {
              max: maxDatetimeValue,
            },
          }}
          onChange={handleDatetimeInputChange}
        />
        <div className="flex flex-row justify-end items-center mt-4 w-full gap-x-2">
          <Button color="neutral" variant="plain" onClick={handleCloseBtnClick}>
            {t("common.cancel")}
          </Button>
          <Button color="primary" onClick={handleSaveBtnClick}>
            {t("common.save")}
          </Button>
        </div>
      </div>
    </>
  );
};

function showChangeLocketCreatedTsDialog(locketId: number) {
  generateDialog(
    {
      className: "change-locket-created-ts-dialog",
      dialogName: "change-locket-created-ts-dialog",
    },
    ChangeLocketCreatedTsDialog,
    {
      locketId,
    },
  );
}

export default showChangeLocketCreatedTsDialog;
