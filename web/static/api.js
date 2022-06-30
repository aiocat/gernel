// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

const normalTitle = "Gernel v0.1.0";
document.title = normalTitle;

const inject = async () =>
  await fetch("/inject", {
    method: "GET",
  });

const execute = async () =>
  await fetch("/execute", {
    method: "POST",
    body: editors[currentTab].getValue(),
  });

const listScripts = async () =>
  await fetch("/scripts", {
    method: "GET",
  });

const getScript = async (name) =>
  await fetch("/scripts", {
    method: "POST",
    body: name,
  });

setInterval(async () => {
  let result = await (await fetch("/injected", { method: "GET" })).text();
  if (result == "1") {
    document.title = `${normalTitle} [INJECTED]`;
  } else {
    document.title = `${normalTitle} [WAITING]`;
  }
}, 100);
