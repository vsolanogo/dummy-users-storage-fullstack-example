import api from "./api";
import { User } from "../models/UsersModels";

export const listUsers = () => {
  return api.get<User[]>("/users").then((res) => res.data);
};

export const createUser = (user: User) => {
  return api.post<User>("/users", user).then((res) => res.data);
};

export const deleteUser = (id: string) => {
  return api.delete(`/users/${id}`);
};

export const patchUser = (user: Partial<User>) => {
  return api.patch<User>(`/users/${user.id}`, user).then((res) => res.data);
};
