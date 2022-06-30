// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

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
