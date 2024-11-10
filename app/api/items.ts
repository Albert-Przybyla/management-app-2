import * as SecureStore from "expo-secure-store";
import { api } from "./api";
import { PagedResponse } from "@/models/pagedResponse";
import { ItemResponse } from "@/models/item/itemResponse.model";

export const fetchItems = async (page: number, pageSize: number): Promise<PagedResponse<ItemResponse>> => {
  const response = await api.get(`/item`, {
    params: { page, pageSize },
  });
  return response.data;
};
