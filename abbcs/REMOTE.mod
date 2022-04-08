MODULE REMOTE
    VAR jointtarget jt;
    VAR robjoint rj;
    VAR rmqmessage msg;
    VAR rmqheader header;
    PROC main()
        TPWrite "Main...";
        WHILE TRUE DO
            RMQReadWait msg \TimeOut:=WAIT_MAX;
            RMQGetMsgHeader msg \Header:=header;
            TPWrite header.datatype;
            IF header.datatype = "robjoint" THEN
                RMQGetMsgData msg, rj;
                TPWrite ValToStr(rj);
                jt := [[0,0,0,0,0,0], [9E9,9E9,9E9,9E9,9E9,9E9]];
                jt.robax := rj;
                MoveAbsJ jt, v1000, fine, tool0;
            ENDIF
        ENDWHILE
    ENDPROC
ENDMODULE