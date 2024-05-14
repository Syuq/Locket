interface CustomizedProfile {
  name: string;
  logoUrl: string;
  description: string;
  locale: Locale;
  appearance: Appearance;
}

interface SystemStatus {
  // System settings
  disablePasswordLogin: boolean;
  disablePublicLockets: boolean;
  maxUploadSizeMiB: number;
  customizedProfile: CustomizedProfile;
  storageServiceId: number;
  localStoragePath: string;
  locketDisplayWithUpdatedTs: boolean;
}

interface SystemSetting {
  name: string;
  value: string;
}
