import { createChannel, createClientFactory, FetchTransport } from "nice-grpc-web";

const channel = createChannel(
  import.meta.env.VITE_API_BASE_URL || window.location.origin,
  FetchTransport({
    credential: "include",
  }),
);

const clientFactory = createClientFactory();
