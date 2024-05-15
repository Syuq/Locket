export const WorkspaceSettingPrefix = "settings/";
export const UserNamePrefix = "users/";
export const LocketNamePrefix = "lockets/";

export const extractLocketIdFromName = (name: string) => {
  return parseInt(name.split(LocketNamePrefix).pop() || "", 10);
};
