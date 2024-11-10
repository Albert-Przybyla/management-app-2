import { Document } from "../document.model";

export interface TransferResponse extends Document {
  source_storage_id: string;
  destination_storage_id: string;
}
