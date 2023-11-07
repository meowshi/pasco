import { useQuery } from "react-query";
import { pick, picks } from "../api/pick";
import CheckIcon from "@mui/icons-material/Check";
import CloseIcon from "@mui/icons-material/Close";
import { Divider } from "@mui/material";
import PersonAddAlt1Icon from "@mui/icons-material/PersonAddAlt1";

export const PickHistory = () => {
  const { isLoading, isError, data } = useQuery("picks", picks, {
    refetchInterval: 5000,
  });

  return (
    <div className="h-[32rem] overflow-y-auto rounded-xl bg-zinc-900">
      {isLoading ? (
        <div>Loading...</div>
      ) : isError ? (
        <div>Error(</div>
      ) : data.data !== null ? (
        data.data.map((pick, i) => (
          <div key={pick.login + pick.event_uuid + pick.picked_at}>
            <div className="m-5 flex items-center justify-between space-x-10">
              <div className="space-y-2 text-lg">
                <div className="flex space-x-2">
                  <div className="text-orange-200">{pick.login}</div>
                  <div>{pick.name}</div>
                  <div>{pick.surname}</div>
                  {pick.with_friends ? <PersonAddAlt1Icon /> : null}
                </div>
                <div>{pick.event_name}</div>
                <div className="text-sm text-gray-500">
                  {new Date(pick.picked_at).toLocaleString("ru-RU")}
                </div>
              </div>
              <div>
                <div className="flex space-x-2 ">
                  {pick.is_list_success ? <CheckIcon /> : <CloseIcon />}
                  <div>Списки</div>
                </div>
                <div className="flex space-x-2">
                  {pick.is_gift_success ? <CheckIcon /> : <CloseIcon />}
                  <div>Подарки</div>
                </div>
                <div className="flex space-x-2">
                  {pick.is_bracelet_success ? <CheckIcon /> : <CloseIcon />}
                  <div>Браслеты</div>
                </div>
              </div>
            </div>
            {data.data.length - 1 !== i ? <Divider /> : null}
          </div>
        ))
      ) : null}
    </div>
  );
};
