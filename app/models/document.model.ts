import { ItemResponse } from "./item/itemResponse.model";

export interface Document {
  id: string;
  notes: string;
  organization_id: string;
  created_at: string;
  updated_at: string;
  items: ItemResponse[];
}
