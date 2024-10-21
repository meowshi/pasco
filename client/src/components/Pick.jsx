import React, { useEffect, useRef, useState } from "react";
import { useQuery, useMutation } from "react-query";
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

  const [autopick, setAutopick] = useState(false);
  const [inputValue, setInputValue] = useState("");
  const [pickCardKey, setPickCardKey] = useState(0);
  const [selectedPrinter, setSelectedPrinter] = useState();

  const keyInputRef = useRef(null);

  const { data: printers } = useQuery(["printers"], getPrinters);
  const pickInfoMutation = useMutation({
    mutationFn: (key) => {
      return pick(key)
    },
    onError: () => {
      enqueueSnackbar("Не удалось найти яндексоида в подарках.", {
        variant: "error",
      });
    },
    retry: 0,
  })

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
                // if (keyInputRef.current.value.length < 6) {
                //   enqueueSnackbar("Неверный rfid.", { variant: "error" });
                //   setInputValue("");
                //   timeoutRunnig = false;
                //   return;
                // }

                const k = rfidToKey(keyInputRef.current.value);
                if (k.length === 0 || k === undefined) {
                  enqueueSnackbar("Неверный rfid.", { variant: "error" });
                  setInputValue("");
                  timeoutRunnig = false;
                  return;
                }
                pickInfoMutation.mutate(k)
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
        {(pickInfoMutation.data !== undefined && pickInfoMutation.data.data !== null) ? (
          <PickCard
            key={pickCardKey}
            selectedPrinter={selectedPrinter}
            pickInfo={pickInfoMutation.data.data}
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
