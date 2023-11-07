import React, { useEffect, useRef, useState } from "react";
import { useQuery } from "react-query";
import {
  ToggleButton,
  ToggleButtonGroup,
  Switch,
  FormControlLabel,
  TextField,
} from "@mui/material";
import PrintIcon from "@mui/icons-material/Print";
import { useSnackbar } from "notistack";
import { getPRinters as getPrinters } from "../api/locker";
import { pick } from "../api/pick";
import { rfidToKey } from "../api/key";
import { PickCard } from "./PickCard";
import { PickHistory } from "./PickHistory";

let timeoutRunnig = false;

const Pick = () => {
  const { enqueueSnackbar } = useSnackbar();

  const [key, setKey] = useState("");
  const [autopick, setAutopick] = useState(false);
  const [inputValue, setInputValue] = useState("");
  const [pickCardKey, setPickCardKey] = useState(0);
  const [selectedPrinter, setSelectedPrinter] = useState();

  const keyInputRef = useRef(null);

  const { data: printers } = useQuery(["printers"], getPrinters);
  const { data: pickInfo, refetch: refetchPick } = useQuery(["pickInfo"], {
    queryFn: () => {
      return pick(key);
    },
    onError: () => {
      enqueueSnackbar("Ошибка при получении пик-информации.", {
        variant: "error",
      });
    },
    refetchOnWindowFocus: false,
    enabled: false,
    retry: 0,
  });

  useEffect(() => {
    if (key.length !== 0) {
      refetchPick();
      setKey("")
    }
  }, [key]);

  return (
    <div className="space-y-10">
      <div className="space-y-2">
        <div className="flex items-center justify-center space-x-14">
          <div className="flex items-center space-x-4">
            <PrintIcon fontSize="large" />
            <ToggleButtonGroup
              value={selectedPrinter}
              onChange={(e) => {
                setSelectedPrinter(e.target.value);
              }}
            >
              {printers !== undefined && printers.data !== null ? (
                printers.data.map((printer) => (
                  <ToggleButton key={printer.id} value={printer.id.toString()}>
                    {printer.name}
                  </ToggleButton>
                ))
              ) : (
                <div>Нема</div>
              )}
            </ToggleButtonGroup>
          </div>
          <TextField
            inputRef={keyInputRef}
            autoFocus
            value={inputValue}
            label="Пикай!"
            onChange={(e) => {
              setInputValue(e.target.value);
              if (timeoutRunnig === true) {
                return;
              }

              timeoutRunnig = true;

              setTimeout(() => {
                if (keyInputRef.current.value.length < 8) {
                  enqueueSnackbar("Неверный rfid.", { variant: "error" });
                  setInputValue("");
                  timeoutRunnig = false;
                  return;
                }

                const k = rfidToKey(keyInputRef.current.value);
                if (k.length === 0 || k === undefined) {
                  enqueueSnackbar("Неверный rfid.", { variant: "error" });
                  setInputValue("");
                  timeoutRunnig = false;
                  return;
                }
                setKey(k);
                setInputValue("");
                setPickCardKey(pickCardKey + 1);
                timeoutRunnig = false;
              }, 250);
            }}
          />
          <FormControlLabel
            control={
              <Switch
                checked={autopick}
                onChange={(e) => setAutopick(e.target.checked)}
              />
            }
            label="Автопик"
          />
        </div>
      </div>
      <div className="flex justify-center">
        {(pickInfo !== undefined && pickInfo.data !== null) ? (
          <PickCard
            key={pickCardKey}
            selectedPrinter={selectedPrinter}
            pickInfo={pickInfo.data}
            autopick={autopick}
          />
        ) : null}
      </div>
      <div className="flex justify-center">
        <PickHistory />
      </div>
    </div>
  );
};

export default Pick;
