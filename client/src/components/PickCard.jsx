import {
  ButtonGroup,
  ToggleButton,
  ToggleButtonGroup,
  Button,
  CircularProgress,
  TextField,
  Checkbox,
  FormControlLabel,
} from "@mui/material";
import { useState, useEffect } from "react";
import { useMutation } from "react-query";
import CheckCircle from "@mui/icons-material/CheckCircle";
import Error from "@mui/icons-material/Error";
import { printBracelet } from "../api/locker";
import { updateRegistration } from "../api/registration";
import { giveGift } from "../api/gift";
import PersonAddAlt1Icon from "@mui/icons-material/PersonAddAlt1";
import { useSnackbar } from "notistack";

function findEvent(events, uuid) {
  if (events === null) {
    return undefined;
  }
  return events.find((value) => value.uuid === uuid);
}

export const PickCard = ({ pickInfo, selectedPrinter, autopick }) => {
  const { enqueueSnackbar } = useSnackbar();

  const [friends, setFriends] = useState(0);

  const [selectedEvent, setSelectedEvent] = useState();
  const [peopleCount, setPeopleCount] = useState(1);

  const listMutation = useMutation({
    mutationFn: (req) => {
      return updateRegistration(req, pickInfo.pick_id);
    },
  });
  const giftMutation = useMutation(giveGift);
  const braceletMutation = useMutation({
    mutationFn: (req) => {
      if (req.printer_id === undefined) {
        enqueueSnackbar("Выбери принтер.", { variant: "error" });
        throw new Error("Принтер не выбран.");
      }

      return printBracelet(req, pickInfo.pick_id);
    },
  });

  useEffect(() => {
    if (pickInfo.events !== null && pickInfo.events.length > 0) {
      setSelectedEvent(pickInfo.events[0].uuid)
      setFriends(pickInfo.events[0].friends)
    }
    if (
      autopick &&
      pickInfo.events !== null &&
      pickInfo.events.length == 1 &&
      !pickInfo.events[0].allowed_friends
    ) {
      console.log("autopick");
      listMutation.mutate({
        event_uuid: pickInfo.events[0].uuid,
        yandexoid_login: pickInfo.login,
        friends: friends,
        status: friends + 1,
        status_cell: pickInfo.events[0].status_cell,
      });

      if (!giftMutation.isSuccess) {
        giftMutation.mutate({
          login: pickInfo.login,
          key: pickInfo.key,
          pickId: pickInfo.pick_id,
        });
      }

      braceletMutation.mutate({
        event_id: pickInfo.events[0].locker_event_id.toString(),
        print_count: friends + 1,
        printer_id: selectedPrinter,
      });
    }
  }, []);


  return (
    <div className="rounded-xl bg-zinc-900">
      <div className="m-5 flex items-center space-x-10">
        <div className="space-y-5">
          <div className="flex space-x-3 text-lg">
            <div className="bold text-orange-200">{pickInfo.login}</div>
            <div>{pickInfo.name}</div>
            <div>{pickInfo.surname}</div>
            {selectedEvent !== undefined &&
            findEvent(pickInfo.events, selectedEvent).friends > 0 ? (
              <div flex items-center className="text-green-600">
                <PersonAddAlt1Icon />
                {" "}
                {findEvent(pickInfo.events, selectedEvent).friends}
              </div>
            ) : null}
          </div>
          <div>
            <ToggleButtonGroup
              className="w-full"
              orientation="vertical"
              value={selectedEvent}
              onChange={(e) => {
                setSelectedEvent(e.target.value);
                setFriends(findEvent(pickInfo.events, e.target.value).friends)
                setPeopleCount(1);
              }}
            >
              {pickInfo.events !== null ? (
                pickInfo.events.map((event) => (
                  <ToggleButton key={event.uuid} value={event.uuid}>
                    {event.name}
                  </ToggleButton>
                ))
              ) : (
                <div className="text-lg text-red-600">
                  Яндексоид никуда не записан!
                </div>
              )}
            </ToggleButtonGroup>
          </div>
          {selectedEvent !== undefined && findEvent(pickInfo.events, selectedEvent).status !== 0 ?
            <div className="text-lg text-red-600">{"Проход был! ("+findEvent(pickInfo.events, selectedEvent).status + " человек/а)"}</div>
            : null
          }
        </div>
        {selectedEvent !== undefined &&
        findEvent(pickInfo.events, selectedEvent).allowed_friends ? (
          <FormControlLabel
            control={
              <div className="flex items-center space-x-2">
                {/* <Checkbox
                  checked={!(peopleCount === 1)}
                  onClick={(e) => {
                    e.target.checked ? setPeopleCount(2) : setPeopleCount(1);
                  }}
                /> */}
                <PersonAddAlt1Icon />
                <TextField
                  value={friends}
                  size="small"
                  style = {{width: 50}}
                  onChange={(e) => {
                    let v = Number(e.target.value)
                    if (isNaN(v)) {
                      setFriends(0)
                      return
                    }
                    if (v >= 100) {
                      setFriends(99)
                      return
                    }
                    setFriends(v)
                  }}
                />
              </div>
            }
          />
        ) : null}
        <div>
          <ButtonGroup orientation="vertical" variant="contained">
            <Button
              onClick={() => {
                const event = findEvent(pickInfo.events, selectedEvent);
                if (event === undefined) {
                  enqueueSnackbar("Выбери событие.", { variant: "error" });
                  throw new Error("Не выбрано событие.");
                }

                listMutation.mutate({
                  event_uuid: event.uuid,
                  yandexoid_login: pickInfo.login,
                  friends: friends,
                  status: friends + 1,
                  status_cell: event.status_cell,
                });
              }}
              startIcon={
                listMutation.isLoading ? (
                  <CircularProgress color="inherit" size={20} />
                ) : listMutation.isSuccess ? (
                  <CheckCircle />
                ) : listMutation.isError ? (
                  <Error />
                ) : null
              }
            >
              Списки
            </Button>
            <Button
              onClick={() => {
                if (!giftMutation.isSuccess) {
                  giftMutation.mutate({
                    login: pickInfo.login,
                    key: pickInfo.key,
                    pickId: pickInfo.pick_id,
                  });
                }
              }}
              startIcon={
                giftMutation.isLoading ? (
                  <CircularProgress color="inherit" size={20} />
                ) : giftMutation.isSuccess ? (
                  <CheckCircle />
                ) : giftMutation.isError ? (
                  <Error />
                ) : null
              }
            >
              Подарки
            </Button>
            <Button
              onClick={() => {
                const event = findEvent(pickInfo.events, selectedEvent);
                if (event === undefined) {
                  enqueueSnackbar("Выбери событие.", { variant: "error" });
                  throw new Error("Не выбрано событие.");
                }

                braceletMutation.mutate({
                  event_id: event.locker_event_id.toString(),
                  print_count: friends + 1,
                  printer_id: selectedPrinter,
                });
              }}
              startIcon={
                braceletMutation.isLoading ? (
                  <CircularProgress color="inherit" size={20} />
                ) : braceletMutation.isSuccess ? (
                  <CheckCircle />
                ) : braceletMutation.isError ? (
                  <Error />
                ) : null
              }
            >
              Браслеты
            </Button>
            <Button
              onClick={() => {
                const event = findEvent(pickInfo.events, selectedEvent);
                if (event === undefined) {
                  enqueueSnackbar("Выбери событие.", { variant: "error" });
                  throw new Error("Не выбрано событие.");
                }

                listMutation.mutate({
                  event_uuid: event.uuid,
                  yandexoid_login: pickInfo.login,
                  friends: friends,
                  status: friends + 1,
                  status_cell: event.status_cell,
                });

                if (!giftMutation.isSuccess) {
                  giftMutation.mutate({
                    login: pickInfo.login,
                    key: pickInfo.key,
                    pickId: pickInfo.pick_id,
                  });
                }

                braceletMutation.mutate({
                  event_id: event.locker_event_id.toString(),
                  print_count: friends + 1,
                  printer_id: selectedPrinter,
                });
              }}
            >
              Пик
            </Button>
          </ButtonGroup>
        </div>
      </div>
    </div>
  );
};
