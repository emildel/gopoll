CREATE OR REPLACE PROCEDURE poll_updates_results(
    IN poll_session text,
    IN pollResultToIncrease int,
    OUT pollAnswers text[],
    OUT pollResults int[]
)
    LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE poll AS p
    SET results[pollResultToIncrease] = results[pollResultToIncrease] + 1
    WHERE p.pollsession = poll_session;

    SELECT p.answers, p.results
    INTO pollAnswers, pollResults
    FROM poll AS p
    WHERE p.pollsession = poll_session;
END;
$$;
