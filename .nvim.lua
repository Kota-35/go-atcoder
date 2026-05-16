vim.lsp.config("gopls", {
  settings = {
    gopls = {
      analyses = {
        unusedparams = false,
      },
      staticcheck = false,
    },
  },
})

vim.lsp.enable("gopls")
