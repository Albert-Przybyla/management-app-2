export interface User {}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface UserResponse {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  organization_id: string;
  role: string;
  created_at: string;
  updated_at: string;
}
