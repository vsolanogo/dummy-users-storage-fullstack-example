import React from "react";
import { useAppDispatch, useAppSelector } from "../store/hooks";
import { useEffect } from "react";
import { Spinner } from "./Spinner";
import AddUsers from "./users/AddUsers";
import DisplayUsers from "./users/DisplayUsers";
import SelectedUserDetails from "./users/SelectedUserDetails";
import { fetchUsers } from "../redux/users/usersSlice";

const UsersManager = () => {
  const dispatch = useAppDispatch();
  const fetchUsersStatus = useAppSelector((state) => {
    return state.users.fetchUsersStatus;
  });

  useEffect(() => {
    dispatch(fetchUsers());
  }, [dispatch]);

  return (
    <div className="container py-8 mx-auto">
      {fetchUsersStatus === "PENDING" ? <Spinner show /> : null}
      {fetchUsersStatus === "SUCCESS" ? (
        <div className="grid grid-cols-12 gap-4 px-4">
          <div className="col-span-4">
            <AddUsers />
          </div>
          <div className="col-span-4">
            <DisplayUsers />
          </div>
          <div className="col-span-4">
            <SelectedUserDetails />
          </div>
        </div>
      ) : null}
      {fetchUsersStatus === "ERROR" ? (
        <p>There was a problem fetching users</p>
      ) : null}
    </div>
  );
};

export default UsersManager;
