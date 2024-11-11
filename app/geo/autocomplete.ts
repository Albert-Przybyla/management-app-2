import { PagedResponse } from "@/models/pagedResponse";
import { ItemResponse } from "@/models/item/itemResponse.model";
import { mapBox } from "./mapbox";

export const fetchItems = async (q: string): Promise<PagedResponse<ItemResponse>> => {
  const response = await mapBox.get(`/search/searchbox/v1/suggest`, {
    params: { q: q, access_token: process.env.EXPO_PUBLIC_MAPBOX_KEY, language: "pl" },
  });
  return response.data;
};
