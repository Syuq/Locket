import { create } from "zustand";
import { combine } from "zustand/middleware";
import { locketServiceClient } from "@/grpcweb";
import { CreateLocketRequest, ListLocketsRequest, Locket } from "@/types/proto/api/v2/locket_service";

interface State {
  locketMapByName: Record<string, Locket>;
}

const getDefaultState = (): State => ({
  locketMapByName: {},
});

export const useLocketStore = create(
  combine(getDefaultState(), (set, get) => ({
    setState: (state: State) => set(state),
    getState: () => get(),
    fetchLockets: async (request: Partial<ListLocketsRequest>) => {
      const { lockets, nextPageToken } = await locketServiceClient.listLockets(request);
      const locketMap = get().locketMapByName;
      for (const locket of lockets) {
        locketMap[locket.name] = locket;
      }
      set({ locketMapByName: locketMap });
      return { lockets, nextPageToken };
    },
    getOrFetchLocketByName: async (name: string, options?: { skipCache?: boolean; skipStore?: boolean }) => {
      const locketMap = get().locketMapByName;
      const locket = locketMap[name];
      if (locket && !options?.skipCache) {
        return locket;
      }

      const res = await locketServiceClient.getLocket({
        name,
      });
      if (!res.locket) {
        throw new Error("Locket not found");
      }

      if (!options?.skipStore) {
        locketMap[name] = res.locket;
        set({ locketMapByName: locketMap });
      }
      return res.locket;
    },
    getLocketByName: (name: string) => {
      return get().locketMapByName[name];
    },
    searchLockets: async (filter: string) => {
      const { lockets } = await locketServiceClient.searchLockets({
        filter,
      });
      const locketMap = get().locketMapByName;
      for (const locket of lockets) {
        locketMap[locket.name] = locket;
      }
      set({ locketMapByName: locketMap });
      return lockets;
    },
    getLocketByUid: (uid: string) => {
      const locketMap = get().locketMapByName;
      return Object.values(locketMap).find((locket) => locket.uid === uid);
    },
    createLocket: async (request: CreateLocketRequest) => {
      const { locket } = await locketServiceClient.createLocket(request);
      if (!locket) {
        throw new Error("Locket not found");
      }

      const locketMap = get().locketMapByName;
      locketMap[locket.name] = locket;
      set({ locketMapByName: locketMap });
      return locket;
    },
    updateLocket: async (update: Partial<Locket>, updateMask: string[]) => {
      const { locket } = await locketServiceClient.updateLocket({
        locket: update,
        updateMask,
      });
      if (!locket) {
        throw new Error("Locket not found");
      }

      const locketMap = get().locketMapByName;
      locketMap[locket.name] = locket;
      set({ locketMapByName: locketMap });
      return locket;
    },
    deleteLocket: async (name: string) => {
      await locketServiceClient.deleteLocket({
        name,
      });

      const locketMap = get().locketMapByName;
      delete locketMap[name];
      set({ locketMapByName: locketMap });
    },
  })),
);

export const useLocketList = () => {
  const locketStore = useLocketStore();
  const lockets = Object.values(locketStore.getState().locketMapByName);

  const reset = () => {
    locketStore.setState({ locketMapByName: {} });
  };

  const size = () => {
    return Object.keys(locketStore.getState().locketMapByName).length;
  };

  return {
    value: lockets,
    reset,
    size,
  };
};
