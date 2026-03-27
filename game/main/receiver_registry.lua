local M = {}

M._registered = {}
M._matched_by_frame = {}

local function key_from_id(id)
    return tostring(id)
end

function M.reset()
    M._registered = {}
    M._matched_by_frame = {}
end

function M.register_receiver(id)
    M._registered[key_from_id(id)] = true
end

function M.unregister_receiver(id)
    M._registered[key_from_id(id)] = nil
end

function M.begin_frame(frame)
    M._matched_by_frame[frame] = {}
    for f, _ in pairs(M._matched_by_frame) do
        if f < frame - 2 then
            M._matched_by_frame[f] = nil
        end
    end
end

function M.mark_result(id, frame, matched)
    -- auto-register in case init order/reset timing differs
    M.register_receiver(id)

    if not matched then return end
    local k = key_from_id(id)
    local bucket = M._matched_by_frame[frame]
    if not bucket then
        bucket = {}
        M._matched_by_frame[frame] = bucket
    end
    bucket[k] = true
end

function M.all_matched(frame)
    local bucket = M._matched_by_frame[frame]
    if not bucket then return false end

    local total = 0
    for _ in pairs(M._registered) do total = total + 1 end
    if total == 0 then return false end

    for k in pairs(M._registered) do
        if not bucket[k] then return false end
    end
    return true
end

return M