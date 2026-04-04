vim.lsp.config('gopls', {
    on_attach = on_attach,
    capabilities = capabilities,
    cmd = {"/Users/ct/.local/share/nvim/mason/bin/gopls"},
    filetypes = {"go", "gomod", "gowork", "gotmpl"},
    settings = {
        gopls = {
            analyses = {
                unusedparams = true, 
                shadow = true,       
            },
            staticcheck = true,      
            gofumpt = true,          
            hints = {
                assignVariableTypes = true,
                compositeLiteralFields = true,
                compositeLiteralTypes = true,
                constantValues = true,
                functionTypeParameters = true,
                parameterNames = true,
                rangeVariableTypes = true,
            },
            completeUnimported = true, 
            usePlaceholders = true,    
        },
    },
})

vim.api.nvim_create_autocmd('FileType', {
    pattern = 'go',
    callback = function(args)
        vim.lsp.enable('gopls', { bufnr = args.buf })
    end,
})