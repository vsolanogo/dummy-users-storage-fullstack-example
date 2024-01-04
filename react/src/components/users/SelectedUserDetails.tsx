import React from "react";
import { useAppSelector } from "../../store/hooks";
import { getSelectedUser } from "../../redux/users/usersSlice";

const SelectedUserDetails = () => {
  const selectedUser = useAppSelector(getSelectedUser);
  return (
    <div>
      <h2 className="font-semibold text-xl mb-4">Selected User Details</h2>
      {selectedUser ? (
        <ul>
          <li>ID: {selectedUser.id}</li>
          <li>Name: {selectedUser.name}</li>
          <li>Email: {selectedUser.email}</li>
        </ul>
      ) : (
        <p>Select a user to see more details</p>
      )}
    </div>
  );
};

export default SelectedUserDetails;
