load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository", "new_git_repository")

go_rules_dependencies()

go_register_toolchains(version = "host")

http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "stream_innovationup_gitlab_innovation_upstream_iu_common_go",
    importpath = "github.com/innovation-upstream/iu-common-go",
    sum = "h1:HjQ32GCPVhn4NaDrhlmoIfPgaKENaqh9wDkZLhTJYSQ=",
    version = "v0.16.4",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:2DntVwHkVopvECVRSlL5PSo9eG+cAkDCuckLubN+rq0=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_go_logr_stdr",
    importpath = "github.com/go-logr/stdr",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_googleapis_enterprise_certificate_proxy",
    importpath = "github.com/googleapis/enterprise-certificate-proxy",
    sum = "h1:RY7tHKZcRlk788d5WSo/e83gOyyy742E8GSs771ySpg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:/nTivgnV/n1CaAeo+ekGexTYUsKEU9jUVkoY5359+3Q=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:CFhNhU7pcD11cuDkQdrE6PQJgv0EXNKNv06jIzbLlCU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:DBi3Jk9XjCJ4pkkLM4NqKgj3ozUL1wq4l+d3/jTGXAI=",
    version = "v1.27.0",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:NKw6PpQi6V1O+KsjuTd+bhip9d0REYu4NevC45vtGp8=",
    version = "v0.12.0",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:IIoXKR7FKrEAQhMTz5hK2wiDz2WNFHS7eVr/L1lE/rM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:AONoTYJviyv1vS4IkvWzq69gEVdvHx35wKXc+e6wjZQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:lmG+O5oaR9xNwaRBwE2XoMhwQHsHql5IoiGr1ptdDwU=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_area120",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:TCMhwWEWhCn8d44/Zs7UCICTWje9j3HuV6nVGMjdpYw=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:3d0LRAU1K6vfqCahhl9fx2oGHcq+s5gftdix4v8Ibrc=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_asset",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:aCrlaLGJWTODJX4G56ZYzJefITKEWNfbjjtHSzWpxW0=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:hhIdCOowsT1GG5eMCIA0OwK6USRuYTou/1ZeNxCSRtA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_automl",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:BMioyXSbg7d7xLibn47cs0elW6RT780IUWr42W8rp2Q=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:g9KO6SkakcYPcc/XjAzeuUrEOXlYPnMpuiaywYaGrmQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_batch",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:1jvEBY55OH4Sd2FxEXQfxGExFWov1A/IaRe+Z5Z71Fw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:w+4kThysgl0JiKshi2MKDCg2NZgOyqOI0wq2eBZyrzA=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:Xkii76HWELHwBtkQVZvqmSo9GTr0O+tIbRNnMcGdlg4=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:pL70vXWn9TitQYXBWTK2abHl2JHLwkFRjYw6VflRqEA=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:tzbR4UHBbgsewMWUD93JHi8EBi/gHBoSAcY1/sThFGk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_channel",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:pNuUlZx0Jb0Ts9P312bmNMuH5IiFWIR4RUtLb70Ke5s=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:TAAmCmAlOJ4uNBu6zwAjwhyl/7fLHHxIEazVhr3QBbQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:UhzHIlgFfMr6luVYVNydw/pl9/U5kgtjCMJHnSvoVws=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:faUiUgXjW8yVZ7XMnKHKm1WE4OldPBUWWfIRN/3z1dc=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_compute",
    importpath = "cloud.google.com/go/compute",
    sum = "h1:hfm2+FfxVmnRlh6LpB7cg1ZNU+5edAHmW679JePztk0=",
    version = "v1.14.0",
)

go_repository(
    name = "com_google_cloud_go_compute_metadata",
    importpath = "cloud.google.com/go/compute/metadata",
    sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
    version = "v0.2.3",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:tTQLI/ZvguUf9Hv+36BkG2+/PeC8Ol1q4pBW+tgCx0A=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_container",
    importpath = "cloud.google.com/go/container",
    sum = "h1:nbEK/59GyDRKKlo1SqpohY1TK8LmJ2XNcvS9Gyom2A0=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:2824iym832ljKdVpCBnpqm5K94YT/uHTVhNF+dRTXPI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:6kZ4RIOW/uT7QWC5SfPfq/G8sYzr/v+UOmOAxy4Z1TE=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:CW3541Fm7KPTyZjJdnX6NtaGXYFn5XbFC5UcjgALKvU=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:vLwowLF2ZB5J5gqiZCzv076lDI/Rd7zYQQFu5XO1PSg=",
    version = "v0.5.0",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:j5m2hjWovTZDTQak4MJeXAR9yN7O+zMfULnjGw/OOLg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:dp8jOF21n/7jwgo/uuA0RN8hvLcKO4q6s/yvwevs2ZM=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:cNxeA2DiWliQGi21kPRqnVeQ5xFhNoEjPRt1400Pm8Y=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_dataproc",
    importpath = "cloud.google.com/go/dataproc",
    sum = "h1:gVOqNmElfa6n/ccG/QDlfurMWwrK3ezvy2b2eDoCmS0=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:gx9jr41ytcA3dXkbbd409euEaWtofCVXYBvJz3iYm18=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:PgIgbhedBtYBU6POGXFMn2uSl9vpqubc3ewTNdcU8Mk=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:kI6dxt8Ml0is/x7YZjLveTvR7YPzXAUD/8wQZ2nH5zA=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:HYHVOkoxQ9bSfNIelSZYNAtUi4CeSrCnROyOsbOqPq8=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:9I4BYeJSVKoSKgjr70fLdRDumqcUeVmHV4fd5f9LR6Y=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:jfq09Fdjtnpnmt/MLyf6A3DM3ynb8B2na0K+vSXvpFM=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_domains",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:pu3JIgC1rswIqi5romW0JgNO6CTUydLYX8zyjiAvO1c=",
    version = "v0.7.0",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:hd6J2n5dBBRuAqnNUEsKWrp6XNPKsaxwwIyzOPZTokk=",
    version = "v0.2.0",
)

go_repository(
    name = "com_google_cloud_go_errorreporting",
    importpath = "cloud.google.com/go/errorreporting",
    sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:b6csrQXCHKQmfo9h3dG/pHyoEh+fQG1Yg78a53LAviY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:AgCqrmMMIcel5WWKkzz5EkCUKC3Rl5LNMMYsS+LvsI0=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:yjKOpzvqtDmL5AXbKttLc8j0hL20kuC1qPdy5HPcxp0=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_functions",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:35tgv1fQOtvKqH/uxJMzX3w6usneJ0zXpsFr9KAVhNE=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_gaming",
    importpath = "cloud.google.com/go/gaming",
    sum = "h1:97OAEQtDazAJD7yh/kvQdSCQuTKdR0O+qWAJBZJ4xiA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:4K+jiv4ocqt1niN8q5Imd8imRoXBHTrdnJVt/uFFxF4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:zAcvDa04tTnGdu6TEZewaLN2tdMtUOJJ7fEceULjguA=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:JTcTaYQRGsVm+qkah7WzHb6e9sf1C0laYdRPn9aN+vg=",
    version = "v0.10.0",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:8F1NhJj8ucNj7lK51UZMtAjSWTgP1zO18XF6vkfiPPU=",
    version = "v0.4.0",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:TGT2oGmO5q3VH6SjcrlgPUWI0njhYv4kywLm6jag0to=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_iam",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:E2osAkZzxI/+8pZcxVLcDtAQx/u+hZXVryUaYQ5O0Kk=",
    version = "v0.8.0",
)

go_repository(
    name = "com_google_cloud_go_iap",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:BGEXovwejOCt1zDk8hXq0bOhhRu9haXKWXXXp2B4wBM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_ids",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:LncHK4HHucb5Du310X8XH9/ICtMwZ2PCfK0ScjWiJoY=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_iot",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:Y9+oZT9jD4GUZzORXTU45XsnQrhxmDT+TFbPil6pRVQ=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_kms",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:OWRZzrPmOZUzurjI2FBGtgY2mB1WaJkqhw6oIwSj0Yg=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_language",
    importpath = "cloud.google.com/go/language",
    sum = "h1:3Wa+IUMamL4JH3Zd3cDZUHpwyqplTACt6UZKRD2eCL4=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:tIqhivE2LMVYkX0BLgG7xL64oNpDaFFI7teunglt1tI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_logging",
    importpath = "cloud.google.com/go/logging",
    sum = "h1:ZBsZK+JG+oCDT+vaxwqF2egKNRjz8soXiS6Xv79benI=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:NjljC+FYPV3uh5/OwWT6pVU+doBqMg2x/rZlE+CamDs=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:3Kdajn6X25yWQFhFCErmKSYTSvkEd3chJROny//F1A0=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_maps",
    importpath = "cloud.google.com/go/maps",
    sum = "h1:kLReRbclTgJefw2fcCbdLPLhPj0U6UUWN10ldG8sdOU=",
    version = "v0.1.0",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:qAJzpxmEX+SeND10Y/4868L5wfZpo4Y3BIEnIieP4dk=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:yLxUzJkZVSH2kPaHut7k+7sbIBFpvSh1LW9qjM2JDjA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:3KcShzqWdqxrDEXIBWpYJpOOrgpDj+HlBi07Grot49Y=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:c9riaGSPQ4dUKWB+M1Fl0N+iLxstMbCktdEwYSPGDvA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:BVdIKaI68bihnXGdCVL89Jsg9kq2kg+II30fjVqo62E=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:mDHA3CDW00imTvC5RW6aMGsD1bH+FtKwZm/52BxaiMg=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:qDEX/3sipg9dS5JYsAY+YvgTjPR63cozzAWop8oZS94=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:AC8RPjNvel3ExgXjO1YOAz+teg9+j+89TNxa7pIZfww=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:7PxOq9VTT7TMib/6dMoWpMvWS2E4dJEvtYzjvBreaec=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:39d6tqvNjd/wsSub1Bn4cEmrYcet5Ur6xpaN+SxOxtY=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:erF5PHqDZb6FeFrUHiYj2JK2BMhsk8CyAg4V4amJ3rE=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:NO0RouqCOM7M2S85Eal6urMSSipWwHU8evzwS+siqUI=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:pKGDPfeZHDybtw48WsnVLjoIPMi9Kw62kUE5TXCLCN4=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:OrwHLSRSZyaiOt3tnY33dsKSedxbMzsXvqB21okItNQ=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:NQklJuOUoz1BPP+Epjw81COx7IISWslkZubz/1i0UN8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:Vz86uiHCtNGm1DeC32HeG2VXmOq5JRYA3VRPf8ZEcSg=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_pubsublite",
    importpath = "cloud.google.com/go/pubsublite",
    sum = "h1:iqrD8vp3giTb7hI1q4TQQGj77cj8zzgmMPsTZtLnprM=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:UqzFfb/WvhwXGDF1eQtdHLrmni+iByZXY4h3w9Kdyv8=",
    version = "v2.5.0",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:6w+WxPf2LmUEqX0YyvfCoYb8aBYOcbIV25Vg6R0FLGw=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:9kMZQGeYfcOD/RtZfcNKGKtoex3DdoB4zRgYU/WaIwE=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_redis",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:/zTwwBKIAD2DEWTrXZp8WD9yD/gntReF/HkPssVYd0U=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:NDao6CHMwEZIaNsdWy+tuvHaavNeGP06o1tgrR0kLvU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:eTzOwB13WrfF0kuzG2ZXCfB3TLunSHBur4s+HFU6uSM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_retail",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:N9fa//ecFUOEPsW/6mJHfcapPV0wBSwIUwpVZB7MQ3o=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_run",
    importpath = "cloud.google.com/go/run",
    sum = "h1:AWPuzU7Xtaj3Jf+QarDWIs6AJ5hM1VFQ+F6Q+VZ6OT4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:K/mxOewgHGeKuATUJNGylT75Mhtjmx1TOkKukATqMT8=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:xE6uXljAC1kCR8iadt9+/blg1fvSbmenlsDN4fT9gqw=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_security",
    importpath = "cloud.google.com/go/security",
    sum = "h1:KSKzzJMyUoMRQzcz7azIgqAUqxo7rmQ5rYvimMhikqg=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:QTVtk/Reqnx2bVIZtJKm1+mpfmwRwymmNvlaFez7fQY=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_servicecontrol",
    importpath = "cloud.google.com/go/servicecontrol",
    sum = "h1:ImIzbOu6y4jL6ob65I++QzvqgFaoAKgHOG+RU9/c4y8=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:f7M8IMcVzO3T425AqlZbP3yLzeipsBHtRza8vVFYMhQ=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_servicemanagement",
    importpath = "cloud.google.com/go/servicemanagement",
    sum = "h1:TpkCO5M7dhKSy1bKUD9o/sSEW/U1Gtx7opA1fsiMx0c=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_serviceusage",
    importpath = "cloud.google.com/go/serviceusage",
    sum = "h1:b0EwJxPJLpavSljMQh0RcdHsUrr5DQ+Nelt/3BAs5ro=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_shell",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:b1LFhFBgKsG252inyhtmsUUZwchqSz3WTvAIf3JFo4g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:NvdTpRwf7DTegbfFdPjAWyD7bOVu0VeMqcvR9aCQCAc=",
    version = "v1.41.0",
)

go_repository(
    name = "com_google_cloud_go_speech",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:yK0ocnFH4Wsf0cMdUyndJQ/hPv02oTJOxzi6AgpBy4s=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:fUe3OydbbvHcAYp07xY+2UpH4AermGbmnm7qdEj3tGE=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_talent",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:MrekAGxLqAeAol4Sc0allOVqUGO8j+Iim8NMvpiD7tM=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:ccPiHgTewxgyAeCWgQWvZvrLmbfQSFABTMAfrSPLPyY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:ztIdKoma1Xob2qm6QwNh4Xi9/e7N3IfvtwG5AcNsj1g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_trace",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:qO9eLn2esajC9sxpqp1YKX37nXC3L4BfGnPS0Cx9dYo=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_translate",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:AOYOH3MspzJ/bH1YXzB+xTE8fMpn3mwhLjugwGXvMPI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_video",
    importpath = "cloud.google.com/go/video",
    sum = "h1:ttlvO4J5c1VGq6FkHqWPD/aH6PfdxujHt+muTJlW1Zk=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:RPFgVVXbI2b5vnrciZjtsUgpNKVtHO/WIyXUhEfuMhA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:TQHxRqvLMi19azwm3qYuDbEzZWmiKJNTpGbkNsfRCik=",
    version = "v2.5.0",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:A2Tl2ZmwMRpvEmhV2ibISY85fmQR+Y5w9a0PlRz5P3s=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_vmwareengine",
    importpath = "cloud.google.com/go/vmwareengine",
    sum = "h1:JMPZaOT/gIUxVlTqSl/QQ32Y2k+r0stNeM1NSqhVP9o=",
    version = "v0.1.0",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:woHXXtnW8b9gLFdWO9HLPalAddBQ9V4LT+1vjKwR3W8=",
    version = "v1.5.0",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:ypSnpGlJnZSXbN9a13PDmAYvVekBLnGKxQ3Q9SMwnYY=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:y7yIFg/h/mO+5Y5aCOtVAnpGUOgqCH5rXQ2Oc8Oq2+g=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:7Chpin9p50NTU8Tb7qk+I11U/IwVXmDhEoSsdccvInE=",
    version = "v1.9.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:n2a8QNdAb0sZNpU9R1ALUXBbY+w51fCQDN+7EdxNBsY=",
    version = "v0.5.0",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:nt+Q6cXKz4MosCSpnbMtqiQ8Oz0pxTef2B4Vca2lvfk=",
    version = "v0.0.0-20221014153046-6fdb5e3db783",
)

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "528927e398f4e290001886894dac17c5c6a2e5548f3fb68004cfb01af901b53a",
    strip_prefix = "protobuf-3.17.3",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.17.3.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
    version = "v0.0.0-20190523083050-ea95bdfd59fc",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:hzAQntlaYRkVSFEfj9OTWlVV1H155FMD8BTKktLv0QI=",
    version = "v0.0.0-20210930031921-04548b0d99d4",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:xvqufLtNVwAhN8NMyWklVgxnWohi+wtMGQMhtxexlm0=",
    version = "v0.10.2-0.20220325020618-49ff273808a1",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:DkWD4oS2D8LGGgTQ6IvwJJXSL5Vp2ffcQg58nFV38Ys=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:DqDEcV5aeaTmdFBePNpYsp3FlcVH/2ISVVM9Qf8PSls=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:1r7pUrabqp18hOBcwBwiTsbnFeTZHV9eER/QT5JVZxY=",
    version = "v0.0.0-20200121045136-8c9f03a8e57e",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:ROPKBNFfQgOUMifHyP+KYbvpjbdoFNs+aK7DXlji0Tw=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:0udJVsspx3VBr5FwtLhQQtuAsVc79tTq0ocGIPAU6qo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
    version = "v0.5.9",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:d8MncMlErDFTwQGBK1xhv026j9kqhvw1Qv9IbWT1VLQ=",
    version = "v3.2.1",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:Jx2lEv4nMccTJE+IIZOVIvk+DjNKlRsW0sm1uBr896U=",
    version = "v0.0.0-20201023163331-3e6fc7fc9c4c",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:IcsPKeInNvYi7eqSaDjiZqDDKu5rsmunY0Y1YupQSSQ=",
    version = "v2.7.0",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:0hERBMJE1eitiLkihrMvRVBYAkpHzc/J3QdDN+dAcgU=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:mV02weKRL81bEnm8A0HT1/CAelMQDBuQIfLw8n+d6xI=",
    version = "v0.0.0-20200824232613-28f6c0f3b639",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:snbPLB8fVfU9iwbbo30TPtbLRzwWu6aJS6Xh4eaaviA=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:qxFzApOv4WsAL965uUPIsXzAKCZxN2p9UqdhFS4ZW10=",
    version = "v0.0.10",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:gQz4mCbXsO+nc9n1hCxHcGA3Zx3Eo+UHZoInFGUIXNM=",
    version = "v0.0.0-20190812154241-14fe0d1b01d4",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:DNtEKRBAAzeS4KyIory52wWHuClNaXJ5x1F7xa4q+5Y=",
    version = "v0.105.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:Wi4dITi+cf9VYp4VH2T9O41w0kCW0uQTELq2Z6tukN0=",
    version = "v1.44.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:4siQRf4zTiAVt/oeH4GureGkApgb2vtPQAtOmhpqQwE=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:IBlRyxgGySXu5VuW0RgGFlTtLukSnNkpDiEOMkQkmpA=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:q+J/Nfr6Qx4RQeu3rJcnN48SNC0qzlYzSeqkPq93VHs=",
    version = "v1.27.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:YOO045NZI9RKfCj1c5A/ZtuuENUc8OAW+gHdGnDgyMQ=",
    version = "v1.27.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "dev_unknwon_clog_v2",
    importpath = "unknwon.dev/clog/v2",
    sum = "h1:jkPdsxux0MC04BT/9NHbT75z4prK92SH10VBNmIpVCc=",
    version = "v2.2.0",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version = "v2.2.2",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
    version = "v0.24.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc",
    sum = "h1:feFCAJIbxME1Kjx+8GLXaJbVynNSRii+PmuYCjmO/P4=",
    version = "v0.10.1",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:1ZAKnNQKwBBxFtww/GwxNUyTf0AxkZzrukO8MeXqe4Y=",
    version = "v1.13.0",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:sW9hgHyX497PP5//NUM7nqfV8D0iDfBApqq7sOh1XR8=",
    version = "v0.109.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:uQ2vjV/sHTsWSqdKeLqmwitzgvjMl7o4IdtHwUDXSJY=",
    version = "v0.0.0-20221227171554-f9683d7f8bef",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:E1eGv1FTqoLIdnBCZufiSHgKjlqG6fKFf6pPWtMTh8U=",
    version = "v1.51.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:d0NfwRgPtno5B1Wa6L2DAG+KivqkdutMf1UhdNx175w=",
    version = "v1.28.1",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:psW17arqaxU48Z5kZ0CQnkZWQJsqcURM6tKiBApRjXI=",
    version = "v0.0.0-20200622213623-75b288015ac9",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
    version = "v0.0.0-20190121172915-509febef88a4",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
    version = "v0.0.0-20190313153728-d0100b6bd8b3",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:6zppjxzCulZykYSLyVDYbneBfbaBIQPYMevg0bEwv2s=",
    version = "v0.6.0-dev.0.20220419223038-86c51ed26bb4",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:L4ZwwTvKW9gr0ZMS1yrHD9GZhIuVjOBBnaKH+SPQK0Q=",
    version = "v0.6.0",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:wsuoTGHzEhffawBOhz5CYhcrV4IdKZbEyZjBMuTp12o=",
    version = "v0.1.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:MUK/U/4lj1t1oPg0HfuXDN/Z1wv31ZJ/YcPiGccS4DU=",
    version = "v0.5.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:4BRB4x83lYWy72KwLD/qYDuTu7q9PjSagHvijDw7cLo=",
    version = "v0.7.0",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:xYY+Bajn2a7VBmTM5GikTmnK8ZuX8YgnQCqZpbBNtmA=",
    version = "v0.1.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
    version = "v0.1.12",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
    version = "v0.0.0-20220907171357-04be3eba64a2",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:vkLuvpK4fmtSCuo60+yC63p7y0BmQ8gm5ZXGuBCJyXg=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_bitly_go_simplejson",
    importpath = "github.com/bitly/go-simplejson",
    sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_bmizerany_assert",
    importpath = "github.com/bmizerany/assert",
    sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
    version = "v0.0.0-20160611221934-b7ed37b82869",
)

go_repository(
    name = "com_github_bugsnag_bugsnag_go",
    importpath = "github.com/bugsnag/bugsnag-go",
    sum = "h1:XnZesafFtcqPryuLUNsWlaTWhdaOcDT66tiorzWVdu4=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_bugsnag_panicwrap",
    importpath = "github.com/bugsnag/panicwrap",
    sum = "h1:2VFfhwhgp270P+wDa/JBt5x7qSm3zOCYoomGX93N8O0=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_datadog_sketches_go",
    importpath = "github.com/DataDog/sketches-go",
    sum = "h1:RtG+76WKgZuz6FIaGsjoPePmadDBkuD/KC6+ZWu78b8=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_gofrs_uuid",
    importpath = "github.com/gofrs/uuid",
    sum = "h1:8K4tyRfvU1CYPgJsveYFQMhpFd/wXNM7iK6rR7UHz84=",
    version = "v3.3.0+incompatible",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:Hsa8mG0dQ46ij8Sl2AYJDUv1oA9/d6Vk+3LG99Oe02g=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_trace",
    importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace",
    sum = "h1:k12Jh/MdFGoOeZc7ciWg7vsqnXlJ7X7T0TXOnSvCZKw=",
    version = "v0.21.0",
)

go_repository(
    name = "com_github_googleinterns_cloud_operations_api_mock",
    importpath = "github.com/googleinterns/cloud-operations-api-mock",
    sum = "h1:eHv/jVY/JNop1xg2J9cBb4EzyMpWZoNCP1BslSAIkOI=",
    version = "v0.0.0-20200709193332-a1e58c29bdd3",
)

go_repository(
    name = "com_github_kardianos_osext",
    importpath = "github.com/kardianos/osext",
    sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
    version = "v0.0.0-20190222173326-2bc1f35cddc0",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:Sy2VLOOg24bipyC29PhuMXYNJrLsxkie8hyI7kUlG9Q=",
    version = "v1.0.0-RC1",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_cncf_xds_go",
    importpath = "github.com/cncf/xds/go",
    sum = "h1:zH8ljVhhq7yC0MIeUL/IviMtY8hx2mK8cN9wEYb8ggw=",
    version = "v0.0.0-20211011173535-cb28da3451f1",
)

go_repository(
    name = "com_github_felixge_httpsnoop",
    importpath = "github.com/felixge/httpsnoop",
    sum = "h1:+nS9g82KMXccJ/wp0zyRW9ZBHFETmMGtkk+2CTTrW4o=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib",
    importpath = "go.opentelemetry.io/contrib",
    sum = "h1:RMJ6GlUVzLYp/zmItxTTdAmr1gnpO/HHMFmvjAhvJQM=",
    version = "v0.21.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    sum = "h1:MUes2rbdXa1ce9mwKYzTyBG0CtqpLT0NgKTFAz8FIDs=",
    version = "v0.39.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
    importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
    sum = "h1:G1vNyNfKknFvrKVC8ga8EYIECy0s5D/QPW4QPRSMhwc=",
    version = "v0.21.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_internal_metric",
    importpath = "go.opentelemetry.io/otel/internal/metric",
    sum = "h1:gZlIBo5O51hZOOZz8vEcuRx/l5dnADadKfpT70AELoo=",
    version = "v0.21.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:t0lgGI+L68QWt3QtOIlqM9gXoxqxWLhZ3R/e5oOAY0Q=",
    version = "v0.36.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_oteltest",
    importpath = "go.opentelemetry.io/otel/oteltest",
    sum = "h1:G685iP3XiskCwk/z0eIabL55XUl2gk0cljhGk9sB0Yk=",
    version = "v1.0.0-RC1",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:CBgRZ6ntv+Amuj1jDsMhZtlAPT6gbyIRdaIzFhfBSdY=",
    version = "v1.13.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:z+mqJhf6ss6BSfSM671tgKyZBFPTTJM+HLxnhPC3wu0=",
    version = "v1.1.10",
)
