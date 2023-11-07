import { Button, CircularProgress, IconButton, TextField } from "@mui/material";
import { useState } from "react";
import { useMutation, useQuery } from "react-query";
import { getEnv, updateSheetTitle } from "../api/env";
import RefreshIcon from "@mui/icons-material/Refresh";

const EnvSettings = () => {
  const [spreadsheetId, setSpreadsheetId] = useState("");
  const [sheetId, setSheetId] = useState("");
  const [sheetTitle, setSheetTitle] = useState("");

  const { isLoading: isSpreadsheetIdLoading, isError: isSpreadsheetIdError } =
    useQuery(
      "spreadsheetId",
      () => {
        return getEnv("GOOGLE_SPREADSHEET_ID");
      },
      { onSuccess: (data) => setSpreadsheetId(data.data.value) },
    );

  const { isLoading: isSheetIdLoading, isError: isSheetIdError } = useQuery(
    "sheetId",
    () => {
      return getEnv("GOOGLE_SHEET_ID");
    },
    { onSuccess: (data) => setSheetId(data.data.value) },
  );

  const {
    isLoading: isSheetTitleLoading,
    isError: isSheetTitleError,
    refetch: refetchSheetTitle,
    isRefetching: isSheetTitleRefetching,
  } = useQuery(
    "sheetTitle",
    () => {
      return getEnv("GOOGLE_SHEET_TITLE");
    },
    {
      onSuccess: (data) => {
        setSheetTitle(data.data.value);
      },
    },
  );

  const mutateTitle = useMutation(updateSheetTitle, {
    onSuccess: refetchSheetTitle,
  });

  return (
    <div className="space-y-5">
      {isSpreadsheetIdLoading ? (
        <div>
          <CircularProgress />
        </div>
      ) : isSpreadsheetIdError ? (
        <div>Error</div>
      ) : (
        <div>
          <TextField
            className="w-full"
            value={spreadsheetId}
            label="Spreadsheet ID"
          />
        </div>
      )}
      {isSheetIdLoading ? (
        <div>
          <CircularProgress />
        </div>
      ) : isSheetIdError ? (
        <div>Error</div>
      ) : (
        <div>
          <TextField className="w-full" value={sheetId} label="Sheet ID" />
        </div>
      )}
      {isSheetTitleLoading || isSheetTitleRefetching ? (
        <div>
          <CircularProgress />
        </div>
      ) : isSheetTitleError ? (
        <div>Error</div>
      ) : (
        <div className="flex items-center space-x-3">
          <TextField
            className="w-full"
            value={sheetTitle}
            label="Sheet Title"
          />
          <IconButton size="large" onClick={() => mutateTitle.mutate()}>
            <RefreshIcon fontSize="inherit" className="h-full" />
          </IconButton>
        </div>
      )}
    </div>
  );
};

export default EnvSettings;
