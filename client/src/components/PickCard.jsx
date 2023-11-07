import {
  ButtonGroup,
  ToggleButton,
  ToggleButtonGroup,
  Button,
  CircularProgress,
  Checkbox,
  FormControlLabel,
} from "@mui/material";
import { useState } from "react";
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

  const [selectedEvent, setSelectedEvent] = useState();
  const [peopleCount, setPeopleCount] = useState(1);

  const markInLists = () => {
    const event = findEvent(pickInfo.events, selectedEvent);
    if (event === undefined) {
      enqueueSnackbar("Выбери событие.", { variant: "error" });
      throw new Error("Не выбрано событие.");
    }

    return updateRegistration(
      {
        event_uuid: event.uuid,
        yandexoid_login: pickInfo.login,
        status: peopleCount,
        status_cell: event.status_cell,
      },
      pickInfo.pick_id,
    );
  };
  const printBracelets = () => {
    const event = findEvent(pickInfo.events, selectedEvent);
    if (event === undefined) {
      enqueueSnackbar("Выбери событие.", { variant: "error" });
      throw new Error("Не выбрано событие.");
    }
    if (selectedPrinter === undefined) {
      enqueueSnackbar("Выбери принтер.", { variant: "error" });
      throw new Error("Принтер не выбран.");
    }

    return printBracelet(
      {
        event_id: event.locker_event_id.toString(),
        print_count: peopleCount,
        printer_id: selectedPrinter,
      },
      pickInfo.pick_id,
    );
  };

  const listMutation = useMutation(markInLists);
  const giftMutation = useMutation(giveGift);
  const braceletMutation = useMutation(printBracelets);

  // useEffect(() => {
  //   if (
  //     autopick &&
  //     pickInfo.events.length == 1 &&
  //     !pickInfo.events[0].allowed_friends
  //   ) {
  //     selectedEvent = pickInfo.events[0];
  //     listMutation.mutate();
  //     giftMutation.mutate({
  //       login: pickInfo.login,
  //       key: pickInfo.key,
  //       pickId: pickInfo.pick_id,
  //     });
  //     braceletMutation.mutate();
  //   }
  // });

  return (
    <div className="rounded-xl bg-zinc-900">
      <div className="m-5 flex items-center space-x-10">
        <div className="space-y-5">
          <div className="flex space-x-3 text-lg">
            <div className="bold text-orange-200">{pickInfo.login}</div>
            <div>{pickInfo.name}</div>
            <div>{pickInfo.surname}</div>
          </div>
          <div>
            <ToggleButtonGroup
              className="w-full"
              orientation="vertical"
              value={selectedEvent}
              onChange={(e) => {
                setSelectedEvent(e.target.value);
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
                <div className="text-lg text-red-500">
                  Яндексоид никуда не записан!
                </div>
              )}
            </ToggleButtonGroup>
          </div>
        </div>
        {selectedEvent !== undefined &&
        findEvent(pickInfo.events, selectedEvent).allowed_friends ? (
          <FormControlLabel
            control={
              <div>
                <Checkbox
                  checked={!(peopleCount === 1)}
                  onClick={(e) => {
                    e.target.checked ? setPeopleCount(2) : setPeopleCount(1);
                  }}
                />
                <PersonAddAlt1Icon />
              </div>
            }
          />
        ) : null}
        <div>
          <ButtonGroup orientation="vertical" variant="contained">
            <Button
              onClick={() => {
                listMutation.mutate();
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
                braceletMutation.mutate();
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
                listMutation.mutate();
                if (!giftMutation.isSuccess) {
                  giftMutation.mutate({
                    login: pickInfo.login,
                    key: pickInfo.key,
                    pickId: pickInfo.pick_id,
                  });
                }
                braceletMutation.mutate();
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
