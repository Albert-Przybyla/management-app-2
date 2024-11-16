import * as SecureStore from "expo-secure-store";
import { api } from "./api";
import { PagedResponse } from "@/models/pagedResponse";
import { TransferResponse } from "@/models/transfer/transferResponse.model";
import { UserResponse } from "@/models/user/user.model";

export const fetchCurrentUser = async (): Promise<UserResponse> => {
  const response = await api.get(`/user/current`);
  return response.data;
};

export const fetchUserList = async (page: number, pageSize: number): Promise<PagedResponse<UserResponse>> => {
  const response = await api.get(`/user`, {
    params: { page, pageSize },
  });
  return response.data;
};
