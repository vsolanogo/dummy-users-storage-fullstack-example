import React from "react";
import { Spinner } from "../Spinner";
import { useAppDispatch, useAppSelector } from "../../store/hooks";
import {
  removeUser,
  selectAllUsers,
  selectUser,
} from "../../redux/users/usersSlice";

const DisplayUsers = () => {
  const dispatch = useAppDispatch();
  const users = useAppSelector(selectAllUsers);
  const deletingUserId = useAppSelector((state) => state.users.deletingUserId);

  return (
    <div>
      <h2 className="font-semibold text-xl mb-4">Users</h2>
      <ul className="space-y-3">
        {users.map((user) => {
          return (
            <li key={user.id} className="space-x-3">
              <button
                className="hover:underline"
                onClick={() => dispatch(selectUser(user.id))}
              >
                {user.email}
              </button>
              {deletingUserId === user.id ? (
                <Spinner show size="sm" />
              ) : (
                <button onClick={() => dispatch(removeUser(user))}>X</button>
              )}
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default DisplayUsers;
