import * as SecureStore from "expo-secure-store";
import { api } from "./api";
import { TOKEN_KEY } from "@env";
import { PagedResponse } from "@/models/pagedResponse";
import { ItemResponse } from "@/models/item/itemResponse.model";

export const fetchItems = async (page: number, pageSize: number): Promise<PagedResponse<ItemResponse>> => {
  const response = await api.get(`/item`, {
    params: { page, pageSize },
  });

  if (response.data.token) {
    await SecureStore.setItemAsync(TOKEN_KEY, response.data.token);
  }

  return response.data;
};
