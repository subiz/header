# Subiz project schedule
```
                            ##
             13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 01 02 03 04
          
    Account  -- -- -- --
      Scope  -- --
     Oauth2  -- --
     Client  -- --

      Email        -- --

   Notifica        -- -- --
   
   Presense              --
 
  Scheduler              -- --

      Convo                 -- ++ ++ -- --
 Subiz chat                 -- ++ ++ -- --
 Email chan                                --
    FB chan                                   -- 
 Viber chan                                      ++   
Channel API                                -- -- ++
     Router                                      ++ ++

    Visitor                                            -- -- --
   Tracking                                                  -- -- --
    Segment                                                     -- -- -- -- --
    Trigger                                                        -- -- -- -- --
      
DashboardUI                                                           ++ ++ -- -- -- -- -- ++ ++ -- -- -- -- -- ++ ++ -- -- -- -- -- ++ ++ -- -- -- -- -- ++ -- -- --
     Widget                                                                                                     ++ ++ -- -- -- -- -- ++ ++ -- -- -- -- -- ++ -- -- --

    Tracing 
 Kubernetes                                                                                                                                                        -- -- --
      Debug                                                                                                                                                        -- -- --

  Billing ?
```


# Rest api
1. clone swagger-ui
2. access http://src/swagger-ui/dist/index.html
3. paste http://src/servicespec/rest/account.yaml to input box

# How to build
## build proto
1. cd /proto
2. ./protoc-go.sh
## build rest
1. npm run rest

## build all
1. ./build.sh

