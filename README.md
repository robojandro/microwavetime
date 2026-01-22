# Handling input like a microwave

This is a quick script meant to simulate how one might accept ambiguous numerical
inputs, validate them, and perform operations as a microwave might resulting in a simulated countdown.

There are a few tricky edge cases that occurred to me while thinking through and coding
up my solution.

I needed to timebox this exercise but was also weary of leaving obvious edge cases unanswered so as a result this took me a bit over 3 hours to complete.

I'm sure there are probably quite a few improvements that could be made as a result.

One glaring omission which I didn't feel I had time for was unit tests.

The decision to support inputs which are improperly formatted quantities of time stems from my own GE microwave supporting this IRL, though here I have opted to convert the input to its corresponding representation of minutes and seconds, which my microwave does not.

![GE microwave accepts 9:99 as input](PXL_20260120_220759449.mp4)

![GE microwave does not convert countdown to minutes and seconds - BUT I will](PXL_20260120_220622936.mp4)

## Example runs

#### Less than a minute is specified
```
$ go run main.go
Enter a cooking time:
33
cooking...
00:33
00:32
00:31
00:30
```

#### Just under 2 minutes is specified in time notation
```
$ go run main.go
Enter a cooking time:
1:59
cooking...
01:59
01:58
01:57
01:56
```

#### An invalid quantity of time

Which we'll convert to 161 seconds or 2 minutes and 41 seconds

```
$ go run main.go
Enter a cooking time:
1:61
cooking...
02:41
02:40
02:39
```

#### Too large a quantity
```
$ go run main.go
Enter a cooking time:
999:99
problem trying to read input: max input is 99:99, got length of 5
exit status 1
```

#### Non-numbers specified

Ignore non-numeric input, but still try to produce a result

```
$ go run main.go
Enter a cooking time:
abc:33
cooking...
00:33
00:32
00:31
00:30
```
