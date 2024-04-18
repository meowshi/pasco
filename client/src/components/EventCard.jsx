import {
  IconButton,
  MenuItem,
  Paper,
  Select,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  InputLabel,
  FormControl,
  Checkbox,
  FormControlLabel,
} from "@mui/material";
import Delete from "@mui/icons-material/Delete";
import { useMutation, useQuery } from "react-query";
import { deleteEvent, getEventLists, updateEvent } from "../api/event";
import { useState } from "react";
import { getLockerEvents } from "../api/locker";
import PersonAddAlt1Icon from "@mui/icons-material/PersonAddAlt1";

const EventCard = ({ props, refetch }) => {
  const { data: lockerEvents } = useQuery(["getLockerEvents"], getLockerEvents);
  const { isLoading, isError, data } = useQuery(
    ["eventLists", props.uuid],
    getEventLists,
  );

  const mutation = useMutation({
    mutationFn: deleteEvent,
    onSuccess: refetch,
  });

  const updateEventMutation = useMutation({
    mutationFn: updateEvent,
  });

  const [isListHidden, setIsListHidden] = useState(true);

  const [allowedFriends, setAllowedFriends] = useState(props.allowed_friends);

  return (
    <div className="flex-1 justify-center">
      <div
        className="grid h-20 grid-cols-2 content-center rounded-xl bg-zinc-900 pr-5"
        onClick={() => {
          setIsListHidden(!isListHidden);
        }}
      >
        <div className="pb-5 pl-5 pt-5 text-xl">{props.name}</div>
        <div className="flex items-center space-x-10 justify-self-end">
          <FormControl className="w-52">
            <InputLabel>Locker</InputLabel>
            <Select
              value={props.locker_event_id}
              label="Locker"
              onClick={(e) => e.stopPropagation()}
              onChange={(e) => {
                props.locker_event_id = Number(e.target.value);
                updateEventMutation.mutate(props);
              }}
            >
              {lockerEvents !== undefined
                ? lockerEvents.data.map((event) => (
                    <MenuItem key={event.id} value={event.id}>
                      {event.description}
                    </MenuItem>
                  ))
                : null}
            </Select>
          </FormControl>
          <FormControlLabel
            className="min-w-max"
            onClick={(e) => e.stopPropagation()}
            control={
              <div>
                <Checkbox
                  checked={props.allowed_friends}
                  onClick={(e) => {
                    e.stopPropagation();
                    setAllowedFriends(e.target.checked);
                    props.allowed_friends = e.target.checked;
                    updateEventMutation.mutate(props);
                  }}
                />
                <PersonAddAlt1Icon />
              </div>
            }
          />
          <IconButton
            onClick={(event) => {
              event.stopPropagation();
              mutation.mutate(props.uuid);
            }}
          >
            <Delete />
          </IconButton>
        </div>
      </div>
      <div>
        {isLoading ? (
          <div>Loading...</div>
        ) : isError ? (
          <div>Error</div>
        ) : data !== null ? (
          <TableContainer hidden={isListHidden} component={Paper}>
            <Table size="small" stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Login</TableCell>
                  <TableCell>Name</TableCell>
                  <TableCell>Surname</TableCell>
                  <TableCell>Friends</TableCell>
                  <TableCell>Status</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {data.map((entry) => (
                  <TableRow key={entry.login}>
                    <TableCell>{entry.login}</TableCell>
                    <TableCell>{entry.name}</TableCell>
                    <TableCell>{entry.surname}</TableCell>
                    <TableCell>{entry.friends}</TableCell>
                    <TableCell>{entry.status}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        ) : null}
      </div>
    </div>
  );
};

export default EventCard;
