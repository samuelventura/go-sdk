MODULE REMOTE
    PROC main()
        TPWrite "Main...";
        WHILE TRUE DO
            handle;
        ENDWHILE
    ENDPROC
    PROC handle()
        VAR bool d_bool;
        VAR rmqmessage msg;
        VAR rmqheader header;
        RMQReadWait msg \TimeOut:=WAIT_MAX;
        RMQGetMsgHeader msg \Header:=header;
        TPWrite header.datatype;
        IF header.datatype = "bool" THEN
            RMQGetMsgData msg, d_bool;
            TPWrite ValToStr(d_bool);
        ENDIF
    ENDPROC    
ENDMODULE