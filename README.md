# HealthSentinel

<p style="text-align:center"><img src="./static/assets/images/common/logo.jpg" style="width:60%;box-shadow:2px 2px 13px 1px #ccc;border-radius:7px"/></p>

(! Beta !)

> English |[简体中文](./README-CN.md)


# What is health Sentinel
Health Sentinel is a small Web program designed to collect health codes for classes in the current epidemic situation in China. It is used to collect three codes:

- health code
- path code
- close code

# How to use it

## step1: install
`
    get github.com/ASWLaunchs/healthSentinel
`

## step2: run

`go run main.go`

## step3: user view

Chinese link: `localhost:22018?lang=zh`
English link: `lcoalhost:22018?lang=en`

## step4: admin view

Chinese link: `localhost:22018/login/?lang=zh`
English link: `localhost:22018/login/?lang=en`


# configuration
All website configured in conf folders that `conf/conf.ini`.
i18n files in `data/i18n`.

# screenshot
<img  style="width:40%;box-shadow:2px 2px 13px 1px #ccc;border-radius:7px" src="./static/assets/images/common/example_admin_en.jpg"/>
<img style="width:40%;box-shadow:2px 2px 13px 1px #ccc;border-radius:7px"  src="./static/assets/images/common/example_admin_index_en.jpg"/>
<img style="width:40%;box-shadow:2px 2px 13px 1px #ccc;border-radius:7px"  src="./static/assets/images/common/example_user_en.jpg"/>

---

# Database dependency

<a href="https://github.com/ostafen/clover"><img style="width:40%;box-shadow:2px 2px 13px 1px #ccc;border-radius:14px;padding:4px 14px"  src="./static/assets/images/common/cloverDB.png"/></a>
