import * as React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import TopBar from "./components/TopBar";
import Pick from "./components/Pick";
import Settings from "./components/Settings";
import ListSettings from "./components/ListSettings";
import EnvSettings from "./components/EnvSettings";
import { QueryClient, QueryClientProvider } from "react-query";
import { SnackbarProvider } from "notistack";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

const queryClient = new QueryClient();

function App() {
  return (
    <SnackbarProvider>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider theme={darkTheme}>
          <CssBaseline />
          <BrowserRouter>
            <TopBar />
            <Routes>
              <Route path="" element={<Pick />} />
              <Route path="settings" element={<Settings />}>
                <Route path="list" element={<ListSettings />} />
                <Route path="env" element={<EnvSettings />} />
              </Route>
            </Routes>
          </BrowserRouter>
        </ThemeProvider>
      </QueryClientProvider>
    </SnackbarProvider>
  );
}

export default App;
