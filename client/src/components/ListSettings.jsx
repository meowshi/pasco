import { useRef, useState } from "react";
import { Button, TextField } from "@mui/material";
import { useSnackbar } from "notistack";
import { useMutation, useQuery } from "react-query";
import EventCard from "./EventCard";
import { getEvents, postEvent } from "../api/event";

const ListSettings = () => {
  const { enqueueSnackbar } = useSnackbar();
  const addEventButton = useRef(null);
  const eventCellInput = useRef(null);

  const [cell, setCell] = useState("");

  const {
    isLoading,
    isError,
    data: events,
    refetch,
  } = useQuery({
    queryKey: "getEvents",
    queryFn: getEvents,
    refetchInterval: 2000,
  });

  const postEventMutation = useMutation({
    mutationFn: postEvent,
    onSuccess: refetch,
    onError: () => {
      enqueueSnackbar("Ошибка при добавлении события.", { variant: "error" });
    },
  });

  return (
    <div className="space-y-3">
      {isLoading ? (
        <div>Loading...</div>
      ) : isError ? (
        <div>Error</div>
      ) : events !== null ? (
        events.map((event) => (
          <EventCard key={event.uuid} refetch={refetch} props={event} />
        ))
      ) : null}
      <div className="flex space-x-5">
        <TextField
          inputRef={eventCellInput}
          value={cell}
          label="Ячейка"
          autoFocus
          onChange={(event) => {
            setCell(event.target.value);
          }}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              addEventButton.current.click();
            }
          }}
        />
        <Button
          ref={addEventButton}
          onClick={(e) => {
            e.stopPropagation();
            e.preventDefault();

            if (cell.length < 2) {
              enqueueSnackbar("Неверно указана ячейка.", { variant: "error" });
              return;
            }

            postEventMutation.mutate({ google_sheet_cell: cell });
            refetch();
            setCell("");
            eventCellInput.current.focus();
          }}
        >
          Добавить событие
        </Button>
      </div>
    </div>
  );
};

export default ListSettings;
