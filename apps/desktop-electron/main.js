const { app, BrowserWindow, shell } = require("electron");
const path = require("path");

function appIndexPath() {
  if (app.isPackaged) {
    return path.join(process.resourcesPath, "web", "static", "index.html");
  }
  return path.resolve(__dirname, "../../web/static/index.html");
}

function createWindow() {
  const window = new BrowserWindow({
    width: 1320,
    height: 900,
    minWidth: 1100,
    minHeight: 760,
    title: "AI NAILS",
    backgroundColor: "#07100f",
    titleBarStyle: process.platform === "darwin" ? "hiddenInset" : "default",
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
      contextIsolation: true,
      nodeIntegration: false,
      sandbox: true,
    },
  });

  window.loadFile(appIndexPath());
  window.webContents.setWindowOpenHandler(({ url }) => {
    shell.openExternal(url);
    return { action: "deny" };
  });
}

app.whenReady().then(() => {
  createWindow();
  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") app.quit();
});
