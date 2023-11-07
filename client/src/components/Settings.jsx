import React from "react";
import { Link, Route, Routes } from "react-router-dom";
import ListSettings from "./ListSettings";
import { Button } from "@mui/material";
import EnvSettings from "./EnvSettings";

const Settings = () => {
  return (
    <div className="flex flex-auto space-x-10">
      <div className="ml-10 flex w-40 flex-col">
        <Button component={Link} to="list">
          Списки
        </Button>
        <Button component={Link} to="env">
          Env
        </Button>
      </div>

      <div className="flex-1">
        <div className="mr-10">
          <Routes>
            <Route path="list" element={<ListSettings />} />
            <Route path="env" element={<EnvSettings />} />
          </Routes>
        </div>
      </div>
    </div>
  );
};

export default Settings;
